import { NgZone } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/observable/throw';

import { environment } from '../../../environments/environment';
import { BaseApiService } from './base-api.service';
import { ProcessingService } from './processing.service';
import { GoogleAnalyticsService } from './google-analytics.service';
import { NotificationsService } from './notifications.service';
import { UserService } from './user.service';
import { IdpUserModel } from '../models/idp-user.model';

import { JustaskingIdpUserModel } from '../models/justasking-idp-user.model';
import { UserModel } from 'app/core/models/user.model';

declare const gapi: any;

@Injectable()
export class IdpAuthenticationService {

  private auth2: any

  constructor(
    private http: Http,
    private userService: UserService,
    private baseApiService: BaseApiService,
    private ngZone: NgZone,
    private router: Router,
    private processingService: ProcessingService,
    private route: ActivatedRoute,
    private googleAnalyticsService: GoogleAnalyticsService,
    private notificationsService: NotificationsService
  ) {
  }

  // JUSTASKING - START //
  registerJustaskingIdpUser(justaskingUser: JustaskingIdpUserModel): Observable<any> {
    let body = justaskingUser;
    var baseUrl = this.baseApiService.getBaseUrl();
    var authResponse = this.http.post(`${baseUrl}justaskinguser`, body)
      .catch(this.baseApiService.handleError)
      .map(response => {
        this.userService.resetUser();
        return response;
      });
    return authResponse;
  }

  initiatePasswordReset(user: UserModel): Observable<any> {
    let body = user;
    var baseUrl = this.baseApiService.getBaseUrl();
    var authResponse = this.http.post(`${baseUrl}user/resetpassword`, body)
      .catch(this.baseApiService.handleError)
      .map(response => {
        this.userService.resetUser();
        return response;
      });
    return authResponse;
  }

  resetPassword(resetCode: string, password: string): Observable<any> {
    let body = { ResetCode: resetCode, Password: password };
    var baseUrl = this.baseApiService.getBaseUrl();
    var authResponse = this.http.put(`${baseUrl}user/password`, body)
      .catch(this.baseApiService.handleError)
      .map(response => {
        this.userService.resetUser();
        return response;
      });
    return authResponse;
  }

  // JUSTASKING - END //

  //GOOGLE - START//
  private loadGoogleAuthApi(): Observable<any> {
    let response = new Observable<any>(observable => {
      gapi.load('auth2', () => {
        this.auth2 = gapi.auth2.init({
          client_id: environment.googleApiClientId,
          cookiepolicy: 'single_host_origin',
          scope: 'profile'
        }).then(function (auth2) {
          observable.next(auth2);
        });
      });
    });
    return response;
  }

  private getGoogleUser(googleUser: any): any {
    let googleIdpUser: any = {};
    let basicProfile: any = googleUser.getBasicProfile();
    let authResponse: any = googleUser.getAuthResponse();

    googleIdpUser.id = basicProfile.getId();
    googleIdpUser.name = basicProfile.getName();
    googleIdpUser.email = basicProfile.getEmail();
    googleIdpUser.imageUrl = basicProfile.getImageUrl();
    googleIdpUser.givenName = basicProfile.getGivenName();
    googleIdpUser.familyName = basicProfile.getFamilyName();
    googleIdpUser.token = authResponse.id_token;

    return googleIdpUser;
  }

  authenticateIdpUser(idpUser: IdpUserModel): Observable<any> {
    var baseUrl = this.baseApiService.getBaseUrl();
    let body = idpUser;
    var authResponse = this.http.post(`${baseUrl}token`, body)
      .catch(this.baseApiService.handleError)
      .map(response => {
        this.userService.resetUser();
        return response;
      });
    return authResponse;
  }

  attachGoogleSignin(element: any, returnUrl: string, plan: string, duration: string) {
    this.loadGoogleAuthApi()
      .catch(error => {
        console.error("Could not load Google Auth Api", error);
        return Observable.throw(error);
      })
      .subscribe(auth2 => {
        auth2.attachClickHandler(element, {}, (googleUser) => {
          this.onGoogleSignInSuccess(googleUser, returnUrl, plan, duration);
        }, this.onGoogleSignInFailure);
      });
  }

  onGoogleSignInFailure = (error) => {
    console.error(error);
  }

  onGoogleSignInSuccess(googleUser: any, returnUrl: string, plan: string, duration: string) {
    this.ngZone.run(
      () => {
        this.processingService.enableProcessingAnimation();
        let idpUser: IdpUserModel = this.getGoogleIdpUser(googleUser);
        this.authenticateIdpUser(idpUser)
          .catch(error => {
            //Let user know to use the Email/Password fields instead
            if (error.status == 409) {
              this.notificationsService.openSnackBarDefault("Sign in using Email address and Password for that account.");
            } else {
              this.notificationsService.openSnackBarDefault("There was an issue logging you in, please try again.");
            }
            this.processingService.disableProcessingAnimation();
            return Observable.throw(error);
          })
          .subscribe(response => {
            let authenticatedToken = (response.json()).Value;
            this.processingService.disableProcessingAnimation();
            this.baseApiService.setToken(authenticatedToken);
            this.baseApiService.setRequestOptionsWithToken(authenticatedToken);
            if (response.status == 201) {
              this.googleAnalyticsService.trackSignUp();
              if (returnUrl && returnUrl !== "/" && returnUrl !== "") {
                this.router.navigate([returnUrl]);
              } else {
                this.router.navigate(['/login/new'], { queryParams: { plan: plan, duration: duration } });
              }
            } else if (response.status == 200) {
              if (returnUrl && returnUrl !== "/" && returnUrl !== "") {
                this.router.navigate([returnUrl]);
              } else {
                this.router.navigate(['/dashboard']);
              }
            }
          });
      }
    )
  }

  getGoogleIdpUser(googleUser: any): any {
    let user: any = this.getGoogleUser(googleUser);

    var idp: IdpUserModel = {
      "IdpName": "Google",
      "IdpData": {
        "name": user.name,
        "email": user.email,
        "imageUrl": user.imageUrl,
        "givenName": user.givenName,
        "familyName": user.familyName,
        "id_token": user.token
      }
    }

    return idp;
  }

  disconnectGoogleAccount() {
    this.loadGoogleAuthApi()
      .catch(error => {
        console.error("Could not disconnect Google's account.", error);
        return Observable.throw(error);
      })
      .subscribe(auth2 => {
        auth2.disconnect();
      });
  }
  //GOOGLE - END//
}
