import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AuthGuard } from '../core/security/auth.guard';

import { LoginComponent } from './login.component';
import { NewUserComponent } from './new-user/new-user.component';
import { ExistingUserComponent } from './existing-user/existing-user.component';
import { SignupUserComponent } from './signup-user/signup-user.component';
import { PasswordResetComponent } from './password-reset/password-reset.component';
import { ForgotPasswordComponent } from './forgot-password/forgot-password.component';
import { GoToComponent } from './go-to/go-to.component';
import { JoinAccountComponent } from './join-account/join-account.component';

const routes: Routes = [
    {
        path: '', component: LoginComponent,
        children: [
            { path: '', component: ExistingUserComponent },
            { path: 'new', component: NewUserComponent },
            { path: 'signup', component: SignupUserComponent },
            { path: 'password-reset', component: PasswordResetComponent },
            { path: 'forgot-password', component: ForgotPasswordComponent },
            { path: 'go-to', component: GoToComponent },
            { path: 'join/:inviteKey', component: JoinAccountComponent, canActivate: [AuthGuard] },
        ]
    },
    { path: '', component: LoginComponent },
    { path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class LoginRoutingModule { }