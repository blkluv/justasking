import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { IdpAuthenticationService } from '../../core/services/idp-authentication.service';
import { ProcessingService } from '../../core/services/processing.service';

import { UserModel } from 'app/core/models/user.model';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.scss']
})
export class ForgotPasswordComponent implements OnInit {
 
  sentPasswordResetEmail:boolean;
  forgotPasswordForm:FormGroup;

  emailFormControl = new FormControl('', [
    Validators.required,
    Validators.email,
  ]);

  constructor(
    private idpAuthenticationService:IdpAuthenticationService,
    private processingService:ProcessingService
  ) { 
    this.sentPasswordResetEmail = false;
  }

  sendResetEmail(user:UserModel){
    this.processingService.enableProcessingAnimation();
    this.idpAuthenticationService.initiatePasswordReset(user)
    .catch(error=>{
      this.processingService.disableProcessingAnimation();
      this.sentPasswordResetEmail = true;
      return Observable.throw(error);
    })
    .subscribe(response=>{
      this.processingService.disableProcessingAnimation();
      this.sentPasswordResetEmail = true;
    });
  }

  ngOnInit() {
    this.forgotPasswordForm = new FormGroup ({
      'Email' : new FormControl('', [
        Validators.required,
        Validators.maxLength(128),
        Validators.pattern(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/),
      ])
    }
  );
  }

}
