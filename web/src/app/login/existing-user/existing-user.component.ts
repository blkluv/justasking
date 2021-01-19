import { Component, NgZone, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { MatDialog } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';

import { TermsAndConditionsDialogComponent } from '../../app-common/terms-and-conditions-dialog/terms-and-conditions-dialog.component';
import { UserModel } from '../../core/models/user.model';
import { IdpUserModel } from '../../core/models/idp-user.model';
import { UserService } from '../../core/services/user.service';
import { BaseApiService } from '../../core/services/base-api.service';
import { IdpAuthenticationService } from '../../core/services/idp-authentication.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';

import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-existing-user',
  templateUrl: './existing-user.component.html',
  styleUrls: ['./existing-user.component.scss'],
})
export class ExistingUserComponent implements OnInit {

  loginFormData: { Email: string; Password: string; };
  returnUrl: string;

  constructor(
    private ngZone: NgZone,
    private dialog: MatDialog,
    private router: Router,
    private route: ActivatedRoute,
    private userService: UserService,
    public processingService: ProcessingService,
    private baseApiService: BaseApiService,
    private notificationsService: NotificationsService,
    private idpAuthenticationService: IdpAuthenticationService
  ) {
    this.loginFormData = {
      Email: "",
      Password: ""
    }
    this.returnUrl = "";
  }

  login() {
    this.processingService.enableProcessingAnimation();
    let idpUser: IdpUserModel = {
      IdpName: "justasking",
      IdpData: { "email": this.loginFormData.Email, "password": this.loginFormData.Password }
    }
    this.idpAuthenticationService.authenticateIdpUser(idpUser)
      .catch(error => {
        //Let user know to use Google instead
        this.notificationsService.openSnackBarDefault("Invalid email or password. Please try again.")
        this.processingService.disableProcessingAnimation();
        return Observable.throw(error);
      })
      .subscribe(response => {
        let authenticatedToken = (response.json()).Value;
        this.processingService.disableProcessingAnimation();
        this.baseApiService.setToken(authenticatedToken);
        this.baseApiService.setRequestOptionsWithToken(authenticatedToken);
        if (this.returnUrl && this.returnUrl !== "/" && this.returnUrl !== "") {
          this.router.navigate([this.returnUrl]);
        } else {
          this.router.navigate(['/dashboard']);
        }
      });
  }

  goToCreateYourAccount() {
    this.router.navigate(["/login/signup"], { queryParams: { returnUrl: this.returnUrl } });
  }

  showTermsAndConditions() {
    let dialogRef = this.dialog.open(TermsAndConditionsDialogComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  ngOnInit() {
    let plan = this.route.snapshot.queryParams['plan'];
    let duration = this.route.snapshot.queryParams['duration'];
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
    let googleSignInButton = document.getElementById('google-login');
    this.idpAuthenticationService.attachGoogleSignin(googleSignInButton, this.returnUrl, plan, duration);
  }
}
