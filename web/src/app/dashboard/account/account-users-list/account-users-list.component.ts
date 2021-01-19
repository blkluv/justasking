import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { NotificationsService } from '../../../core/services/notifications.service';
import { ProcessingService } from '../../../core/services/processing.service';
import { AccountService } from '../../../core/services/account.service';
import { UserService } from '../../../core/services/user.service';

import { AccountUserModel } from '../../../core/models/account-user.model';
import { AccountUserInviteModel } from '../../../core/models/account-user-invite.model';
import { AccountModel } from '../../../core/models/account.model';

@Component({
  selector: 'app-account-users-list',
  templateUrl: './account-users-list.component.html',
  styleUrls: ['./account-users-list.component.scss']
})
export class AccountUsersListComponent implements OnInit {

  accountUsers: AccountUserModel[];
  inviteUserForm: FormGroup;
  gettingUsers: boolean;
  assignableRoles: any[];

  constructor(
    public userService: UserService,
    private accountService: AccountService,
    private notificationsService: NotificationsService,
    public processingService: ProcessingService
  ) {
    this.gettingUsers = false;
    this.accountUsers = [];
  }

  inviteUser(inviteUserFormData: any) {
    let invitedUser: AccountUserModel = {
      Email: inviteUserFormData.Email,
      RoleId: inviteUserFormData.Role.RoleId,
      RoleName: inviteUserFormData.Role.RoleName,
    };
    let userIsInAccount = this.isUserInAccount(invitedUser, this.accountUsers);
    if (!userIsInAccount) {
      this.processingService.enableProcessingAnimation();
      this.accountService.inviteUserToAccount(invitedUser)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          if (error.status == 409) {
            this.inviteUserForm.reset();
            this.notificationsService.openSnackBarDefault("User is already on this account.");
          } else if (error.status == 403) {
            this.inviteUserForm.reset();
            this.notificationsService.openSnackBarDefault("User limit reached for this account.");
          }
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.inviteUserForm.reset();
          this.processingService.disableProcessingAnimation();
          invitedUser.IsPending = true;
          invitedUser.IsActive = false;
          this.accountUsers.push(invitedUser)
          this.notificationsService.openSnackBarDefault("User has been invited!");
        });
    } else {
      this.inviteUserForm.reset();
      this.notificationsService.openSnackBarDefault("User is already on this account.");
    }
  }

  private isUserInAccount(invitedUser: AccountUserModel, accountUsers: AccountUserModel[]) {
    let isUserInAccount = false;
    let invitedUserEmail = invitedUser.Email && invitedUser.Email.toLowerCase().trim();
    if (accountUsers && accountUsers.length > 0) {
      accountUsers.forEach(currentUser => {
        let currentInvitedUserEmail = currentUser.Email && currentUser.Email.toLowerCase().trim();
        if (currentInvitedUserEmail == invitedUserEmail) {
          isUserInAccount = true;
          return;
        }
      });
    }
    return isUserInAccount;
  }

  private removeUserFromAccountList(accountUser:AccountUserModel){
    var index = this.accountUsers.indexOf(accountUser);
    this.accountUsers.splice(index, 1);
  }

  ngOnInit() {
    this.gettingUsers = true;
    this.inviteUserForm = new FormGroup({
      'Email': new FormControl('', [
        Validators.required,
        Validators.maxLength(128),
        Validators.pattern(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/),
      ]),
      'Role': new FormControl('', [
        Validators.required
      ])
    }); 
    this.userService.getUser()
      .catch(error => {
        this.gettingUsers = false;
        return Observable.throw(error);
      })
      .subscribe((user) => {
        this.assignableRoles = this.accountService.getAssignableRoles();
        this.accountService.getAllUsersForAccount({})
          .catch(error => {
            this.gettingUsers = false;
            return Observable.throw(error);
          })
          .subscribe((users: AccountModel[]) => {
            this.gettingUsers = false;
            this.accountUsers = users;
          });
      });
  }

}