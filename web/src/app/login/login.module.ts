import { NgModule } from '@angular/core';

import {RecaptchaModule} from 'ng-recaptcha';

import { AppCommonModule } from '../app-common/app-common.module';

import { LoginComponent } from './login.component';
import { LoginRoutingModule } from './login.routing.module';
import { NewUserComponent } from './new-user/new-user.component';
import { SignupUserComponent } from './signup-user/signup-user.component';
import { ExistingUserComponent } from './existing-user/existing-user.component';
import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';
import { PasswordResetComponent } from './password-reset/password-reset.component';
import { GoToComponent } from './go-to/go-to.component';
import { JoinAccountComponent } from './join-account/join-account.component';

@NgModule({
  imports: [
    LoginRoutingModule,
    AppCommonModule,
    RecaptchaModule.forRoot()
  ],
  declarations: [
    GoToComponent,
    LoginComponent, 
    NewUserComponent, 
    SignupUserComponent,
    ExistingUserComponent,  
    ForgotPasswordComponent,
    PasswordResetComponent,
    JoinAccountComponent
  ]
})
export class LoginModule { }
