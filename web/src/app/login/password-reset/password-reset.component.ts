import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Params } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { BaseApiService } from '../../core/services/base-api.service';
import { IdpAuthenticationService } from '../../core/services/idp-authentication.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { JustaskingIdpUserModel } from '../../core/models/justasking-idp-user.model';
import { PasswordValidator } from 'app/core/validators/password.validator';

@Component({
  selector: 'app-password-reset',
  templateUrl: './password-reset.component.html',
  styleUrls: ['./password-reset.component.scss']
})
export class PasswordResetComponent implements OnInit {

  resetPasswordToken: string;
  newPasswordSaved: boolean;
  newPasswordForm: FormGroup;

  constructor(
    private route: ActivatedRoute,
    private idpAuthenticationService: IdpAuthenticationService,
    private baseApiService: BaseApiService,
    private processingService: ProcessingService,
    private notificationsService: NotificationsService
  ) {
    this.newPasswordSaved = false;
  }

  saveNewPassword(user: JustaskingIdpUserModel) {
    this.processingService.enableProcessingAnimation();
    this.idpAuthenticationService.resetPassword(this.resetPasswordToken, user.Password)
      .catch(error => {
        if (error.status == 410) {
          this.notificationsService.openSnackBarDefault("This reset link is no longer valid, please request a new link to reset your password.");
        }
        this.processingService.disableProcessingAnimation();
        return Observable.throw(error);
      })
      .subscribe(response => {
        let authenticatedToken = response.json();
        authenticatedToken = authenticatedToken && authenticatedToken.Value;
        this.baseApiService.setToken(authenticatedToken);
        this.baseApiService.setRequestOptionsWithToken(authenticatedToken);
        this.processingService.disableProcessingAnimation();
        this.newPasswordSaved = true;
      });
  }

  ngOnInit() {
    this.resetPasswordToken = this.route.snapshot.queryParams['t'] || '';
    this.newPasswordForm = new FormGroup({
      'Password': new FormControl('', [
        Validators.required,
        Validators.minLength(12),
        Validators.maxLength(32)
      ]),
      'ConfirmPassword': new FormControl('', [
        Validators.required
      ])
    },
      {
        validators: [
          PasswordValidator.MatchPassword
        ]
      }
    );
  }

}
