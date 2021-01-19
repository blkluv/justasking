import { Component, OnInit } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { AccountUserModel } from '../../core/models/account-user.model';

import { UserService } from '../../core/services/user.service';
import { ProcessingService } from '../../core/services/processing.service';
import { SidenavService } from '../../core/services/sidenav.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { AccountService } from '../../core/services/account.service';
import { BaseApiService } from '../../core/services/base-api.service';

@Component({
  selector: 'app-dashboard-top-nav',
  templateUrl: './dashboard-top-nav.component.html',
  styleUrls: ['./dashboard-top-nav.component.scss']
})
export class DashboardTopNavComponent implements OnInit {

  constructor(
    public userService: UserService,
    public processingService: ProcessingService,
    private notificationsService: NotificationsService,
    private sidenavService: SidenavService,
    public accountService: AccountService,
    private baseApiService: BaseApiService,
  ) {
  }

  toggleSidenav() {
    this.sidenavService.toggle();
  }

  openSidenav() {
    this.sidenavService.open();
  }

  closeSidenav() {
    this.sidenavService.close();
  }

  changeCurrentAccount(account: AccountUserModel) {
    this.processingService.enableProcessingAnimation();
    this.accountService.changeCurrentAccount(account.AccountId)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError(`Could not change account. Please try again.`);
        return Observable.throw(error);
      })
      .subscribe(newToken => {
        this.baseApiService.setToken(newToken);
        this.baseApiService.updatedToken(newToken);
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarDefault(`Changed current account to : ${account.AccountName}`);
        window.location.reload();
        //ENABLE ONCE YOU FIGURE OUT HOW TO RELOAD COMPONENTS WITHOUT REFRESHING SITE
        // this.userService.getUserFromApi()
        //   .catch(error => {
        //     this.processingService.disableProcessingAnimation();
        //     this.notificationsService.openSnackBarError(`Could not change account. Please try again.`);
        //     return Observable.throw(error);
        //   })
        //   .subscribe(() => {
        //     this.processingService.disableProcessingAnimation();
        //     this.notificationsService.openSnackBarDefault(`Changed current account to : ${account.AccountName}`);
        //     window.location.reload();
        //   })
      });
  }

  ngOnInit() {
    this.accountService.getAccountsForUsers()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe();
  }

}
