import { NgModule } from '@angular/core';

import { AppCommonModule } from '../../app-common/app-common.module';
import { AccountRoutingModule } from './account.routing.module';

import { AccountComponent } from './account.component';
import { AccountUpgradeComponent } from './account-upgrade/account-upgrade.component';
import { AccountUpgradeCustomComponent } from './account-upgrade-custom/account-upgrade-custom.component';
import { AccountUserProfileComponent } from './account-user-profile/account-user-profile.component';
import { AccountUsersListComponent } from './account-users-list/account-users-list.component';
import { AccountDetailsComponent } from './account-details/account-details.component';
import { AccountUserListItemComponent } from './account-user-list-item/account-user-list-item.component';
import { AlreadyUpgradedComponent } from './already-upgraded/already-upgraded.component';

@NgModule({
  imports: [
    AccountRoutingModule,
    AppCommonModule
  ],
  declarations: [
    AccountComponent,
    AccountUpgradeComponent,
    AccountUpgradeCustomComponent,
    AccountUserProfileComponent,
    AccountUsersListComponent,
    AccountDetailsComponent,
    AccountUserListItemComponent,
    AlreadyUpgradedComponent
  ] 
})
export class AccountModule { }
