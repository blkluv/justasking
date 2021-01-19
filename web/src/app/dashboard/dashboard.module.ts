import { NgModule } from '@angular/core';
import { AppCommonModule } from '../app-common/app-common.module';

import { DashboardComponent } from './dashboard.component';
import { DashboardTopNavComponent } from './dashboard-top-nav/dashboard-top-nav.component';
import { DashboardRoutingModule } from './dashboard.routing.module';

@NgModule({
  imports: [
    DashboardRoutingModule,
    AppCommonModule,
  ],
  declarations: [
    DashboardComponent,
    DashboardTopNavComponent
  ]
})
export class DashboardModule { }
