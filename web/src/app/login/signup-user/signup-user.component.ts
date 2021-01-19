import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { MatSelectChange, MatSlideToggleChange } from '@angular/material';
import { Location } from '@angular/common';
import { MatDialog } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/observable/throw';

import { TermsAndConditionsDialogComponent } from '../../app-common/terms-and-conditions-dialog/terms-and-conditions-dialog.component';

import { BaseApiService } from '../../core/services/base-api.service';
import { AccountService } from '../../core/services/account.service';
import { PlanService } from '../../core/services/plan.service';
import { StripeService } from '../../core/services/stripe.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { IdpAuthenticationService } from '../../core/services/idp-authentication.service';
import { GoogleAnalyticsService } from '../../core/services/google-analytics.service';

import { JustaskingIdpUserModel } from '../../core/models/justasking-idp-user.model';
import { UserModel } from 'app/core/models/user.model';
import { PasswordValidator } from 'app/core/validators/password.validator';

import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-signup-user',
  templateUrl: './signup-user.component.html',
  styleUrls: ['./signup-user.component.scss']
})
export class SignupUserComponent implements OnInit {
 
  confirmedPassword:string;
  registeringUser:boolean; 
  signUpForm:FormGroup;
  reCaptchaSiteKey:string;
  captchaResponse:string;
  returnUrl:string;

  constructor(
    private dialog:MatDialog,
    private baseApiService: BaseApiService,
    private accountService: AccountService,
    private planService: PlanService,
    private stripeService: StripeService,
    private router: Router,
    private route: ActivatedRoute, 
    private location: Location,
    private processingService:ProcessingService,
    private idpAuthenticationService:IdpAuthenticationService,
    private notificationsService:NotificationsService,
    private googleAnalyticsService:GoogleAnalyticsService 
  ) { 
    this.reCaptchaSiteKey = environment.reCaptchaSiteKey;
    this.confirmedPassword = "";
    this.returnUrl = "";
    this.registeringUser = false;
  }
  
  showTermsAndConditions(){
    let dialogRef = this.dialog.open(TermsAndConditionsDialogComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  registerUser(newUser:JustaskingIdpUserModel){
    this.registeringUser = true;
    this.processingService.enableProcessingAnimation();

    //cleaning up user entered data
    newUser.Email = newUser.Email.trim();
    newUser.GivenName = newUser.GivenName.trim();
    newUser.FamilyName = newUser.FamilyName.trim();
    newUser.PhoneNumber = newUser.PhoneNumber.trim();
    newUser.Name = `${newUser.GivenName} ${newUser.FamilyName}`;
    newUser.CaptchaToken = this.captchaResponse; 
 
    this.idpAuthenticationService.registerJustaskingIdpUser(newUser)
    .catch(error=>{
      if(error.status == 409){
        this.notificationsService.openSnackBarDefault("An account with that email already exists. Please sign in.")
      }else{
        this.notificationsService.openSnackBarDefault("There was an issue creating your account. Continue with Google if problem persists.")
      }
      this.registeringUser = false;
      this.processingService.disableProcessingAnimation();
      return Observable.throw(error);
    })
    .subscribe(response=>{
      this.registeringUser = false;
      let authenticatedToken = response.json();
      this.baseApiService.setToken(authenticatedToken);
      this.baseApiService.setRequestOptionsWithToken(authenticatedToken);
      this.processingService.disableProcessingAnimation();
      this.googleAnalyticsService.trackSignUp();
      if(this.returnUrl && this.returnUrl !=="/" && this.returnUrl !==""){
        this.router.navigate([this.returnUrl]);
      } else{
        this.router.navigate(['/dashboard/new'], {queryParams: {welcome:true}});
      }
    });
  }

  resolved(captchaResponse: string) {
    this.captchaResponse = captchaResponse;
}

  ngOnInit() { 
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
    let preselectedPlanName = this.route.snapshot.queryParams['plan'] || 'BASIC';
    let isYearPlan = this.route.snapshot.queryParams['yearPlan'] == 'true';
    this.signUpForm = new FormGroup ({
      'GivenName' : new FormControl('', [
        Validators.required,
        Validators.maxLength(128)
      ]),
      'FamilyName' : new FormControl('', [
        Validators.required,
        Validators.maxLength(128)
      ]),
      'Email' : new FormControl('', [
        Validators.required,
        Validators.maxLength(128),
        Validators.pattern(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/),
      ]),
      'Password' : new FormControl('', [
        Validators.required,
        Validators.minLength(12),
        Validators.maxLength(32)
      ]),
      'ConfirmPassword' : new FormControl('', [
        Validators.required
      ]),
      'PhoneNumber' : new FormControl('', [ 
        Validators.maxLength(25)
      ]),
    },
    {
      validators: [
        PasswordValidator.MatchPassword
      ]
    }
  );
  }
}
