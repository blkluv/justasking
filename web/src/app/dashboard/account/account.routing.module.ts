import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AccountComponent } from './account.component';
import { AccountUpgradeComponent } from './account-upgrade/account-upgrade.component';
import { AccountUpgradeCustomComponent } from './account-upgrade-custom/account-upgrade-custom.component';

const routes: Routes = [
{ path: '', component: AccountComponent },
{ path: 'upgrade', component: AccountUpgradeComponent },
{ path: 'upgrade/custom/:licenseCode', component: AccountUpgradeCustomComponent },
{ path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class AccountRoutingModule {}