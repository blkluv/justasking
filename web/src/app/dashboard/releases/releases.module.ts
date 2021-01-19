import { NgModule } from '@angular/core';

import { AppCommonModule } from '../../app-common/app-common.module';
import { ReleasesRoutingModule } from './releases.routing.module';

import { ReleasesComponent } from './releases.component';

@NgModule({
  imports: [
    ReleasesRoutingModule,
    AppCommonModule,
  ],
  declarations: [ReleasesComponent]
})
export class ReleasesModule { }
