import { NgModule } from '@angular/core';

import { AppCommonModule } from '../app-common/app-common.module';
import { LandingRoutingModule } from './landing.routing.module'
import { LandingComponent } from './landing.component';
import { PricingComponent } from './pricing/pricing.component';
import { GetstartedComponent } from './getstarted/getstarted.component';

@NgModule({
  imports: [
    LandingRoutingModule,
    AppCommonModule,
  ],
  declarations: [
      LandingComponent,
      PricingComponent,
      GetstartedComponent
  ],
  entryComponents: [
  ]
})
export class LandingModule { }
