import { NgModule } from '@angular/core';

import { AppCommonModule } from '../../app-common/app-common.module';
import { BoxesComponent } from './boxes.component';
import { AccountUpgradedDialogComponent } from './account-upgraded-dialog/account-upgraded-dialog.component';
import { BoxesRoutingModule } from './boxes.routing.module';

@NgModule({
  imports: [
    BoxesRoutingModule,
    AppCommonModule,
  ],
  declarations: [
    BoxesComponent,
    AccountUpgradedDialogComponent
  ],
  entryComponents: [
    AccountUpgradedDialogComponent
  ]
})
export class BoxesModule { }
