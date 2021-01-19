import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ConfirmationDialogComponent } from '../../../app-common/confirmation-dialog/confirmation-dialog.component';

import { NotificationsService } from '../../../core/services/notifications.service';
import { ProcessingService } from '../../../core/services/processing.service';
import { AccountService } from '../../../core/services/account.service';
import { BaseApiService } from '../../../core/services/base-api.service';

import { AccountUserInviteModel } from '../../../core/models/account-user-invite.model';
import { AccountUserModel } from '../../../core/models/account-user.model';
import { UserModel } from '../../../core/models/user.model';
import { RoleModel } from 'app/core/models/role.model';
import { Roles } from 'app/core/constants/roles.constants';

@Component({
  selector: 'app-account-user-list-item',
  templateUrl: './account-user-list-item.component.html',
  styleUrls: ['./account-user-list-item.component.scss']
})
export class AccountUserListItemComponent implements OnInit {

  assignableRoles: any[];
  @Input('accountUser') accountUser: AccountUserModel;
  @Input('currentUser') currentUser: UserModel;
  @Output() accountUserRemoved: EventEmitter<AccountUserModel> = new EventEmitter();
  canManageUserAccount: boolean;
  initialAccountUserRole: RoleModel;
  accountUserRole: RoleModel;

  constructor(
    private dialog: MatDialog,
    private accountService: AccountService,
    private notificationsService: NotificationsService,
    private baseApiService: BaseApiService,
    public processingService: ProcessingService) {
  }

  confirmCancelAccountInvitation(accountUser: AccountUserModel) {
    let dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        Header: "Cancel invitation",
        Texts: ["Are you sure you want to cancel this invitation?"],
        ConfirmText: "Yes, cancel the invitation",
        ConfirmAction: () => { return this.cancelAccountInvitation(accountUser); }
      }
    });
  }

  cancelAccountInvitation(accountUser: AccountUserModel): Observable<any> {
    let observableResponse = new Observable<any>(observable => {
      this.processingService.enableProcessingAnimation();
      let accountInvitation: AccountUserInviteModel = {
        AccountId: this.currentUser.Account.Id,
        Email: accountUser.Email
      };
      this.accountService.cancelAccountInvitation(accountInvitation)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          observable.next(false);
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.accountUserRemoved.emit(accountUser);
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault("User invitation has been cancelled.")
          observable.next(true);
        });
    });
    return observableResponse;
  }

  confirmRemoveUserFromAccount(accountUser: AccountUserModel) {
    let dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        Header: "Remove user",
        Texts: ["Are you sure you want to remove this user?"],
        ConfirmText: "Yes, remove user",
        ConfirmAction: () => { return this.removeUserFromAccount(accountUser); }
      }
    });
  }

  removeUserFromAccount(accountUser: AccountUserModel): Observable<any> {
    let observableResponse = new Observable<any>(observable => {
      accountUser.IsActive = false;
      this.processingService.enableProcessingAnimation();
      this.accountService.removeAccountUser(accountUser)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault("Could not remove user from account.");
          observable.next(false);
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.accountUserRemoved.emit(accountUser);
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault("Removed user from account.");
          observable.next(true);
        });
    });
    return observableResponse;
  }

  confirmUpdateAccountUserRole(accountUser: AccountUserModel) {
    let dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        Header: "Update user role",
        Texts: [`Are you sure you want to update the role for <b>${accountUser.Email}</b> to <b>${this.accountUserRole.RoleName}</b>?`, this.accountUserRole.RoleDescription],
        ConfirmText: "Yes, update role",
        ConfirmAction: () => { return this.updateAccountUserRole(); }
      }
    });
    dialogRef.beforeClose()
      .subscribe(response => {
        this.accountUserRole = this.initialAccountUserRole;
      });
  }

  updateAccountUserRole(): Observable<any> {
    let observableResponse = new Observable<any>(observable => {
      this.processingService.enableProcessingAnimation();
      this.accountUser.RoleId = this.accountUserRole.RoleId;
      this.accountUser.RoleName = this.accountUserRole.RoleName;
      this.accountService.updateAccountUser(this.accountUser)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault("Could not update user role!");
          this.accountUserRole = this.initialAccountUserRole;
          observable.next(false);
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault("Updated user role!");
          this.initialAccountUserRole = this.accountUserRole;
          observable.next(true);
        });
    });
    return observableResponse;
  }

  confirmTransferAccountOwnership(accountUser: AccountUserModel) {
    let dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        Header: "Transfer account ownership",
        Texts: [
          `Are you sure you want to make <b>${accountUser.Email}</b> the owner of this account?`,
          `You will become an <b>Admin</b> and no longer have access to: <ul><li>Update account details</li><li>Manage account billing</li></ul>`,
        ],
        ConfirmText: "Yes, transfer account ownership",
        ConfirmAction: () => { return this.transferAccountOwnership(accountUser); }
      }
    });
  }

  transferAccountOwnership(accountUser: AccountUserModel): Observable<any> {
    let observableResponse = new Observable<any>(observable => {
      this.processingService.enableProcessingAnimation();
      this.accountService.transferAccountOwnership(accountUser)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault("Could not transfer account ownership.");
          observable.next(false);
          return Observable.throw(error);
        })
        .subscribe(newToken => {
          this.baseApiService.setToken(newToken);
          this.baseApiService.updatedToken(newToken);
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarDefault(`Account ownership transfered to ${accountUser.Email}`);
          window.location.reload();
          observable.next(true);
        });
    });
    return observableResponse;
  }

  canManageAccount(accountUser: AccountUserModel, currentUser: UserModel): boolean {
    let canChangeRole = false;

    if (currentUser.RolePermissions.ManageUsers) {
      canChangeRole = true;
    }
    if (accountUser.Email === currentUser.Email) {
      canChangeRole = false;
    } else if (accountUser.IsPending) {
      canChangeRole = false;
    } else if (accountUser.RoleName.toLowerCase() === 'owner' && currentUser.RolePermissions.ManageUsers && !(currentUser.RolePermissions.ManageOwners)) {
      canChangeRole = false;
    }

    return canChangeRole;
  }

  initializeAccountUserRole(assignableRoles: RoleModel[], accountUser: AccountUserModel) {
    assignableRoles.forEach((role: RoleModel) => {
      if (role.RoleId === accountUser.RoleId) {
        this.accountUserRole = role;
        this.initialAccountUserRole = this.accountUserRole;
        return;
      }
    });
  }

  resendUserAccountInvitation(accountUser: AccountUserModel){
    this.processingService.enableProcessingAnimation();
    this.accountService.resendUserAccountInvitation(accountUser)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarDefault("Could not send account invitation.");
        return Observable.throw(error);
      })
      .subscribe(newToken => {
        this.baseApiService.setToken(newToken);
        this.baseApiService.updatedToken(newToken);
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarDefault(`Account invitation sent to ${accountUser.Email}`);
      });
  }

  ngOnInit() {
    this.canManageUserAccount = this.canManageAccount(this.accountUser, this.currentUser);
    this.assignableRoles = this.accountService.getAssignableRoles();
    this.initializeAccountUserRole(this.assignableRoles, this.accountUser);
  }
}
