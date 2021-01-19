import { NgModule } from '@angular/core';

import { AppCommonModule } from '../../app-common/app-common.module';
import { FeedbackRoutingModule } from './feedback.routing.module';

import { FeedbackComponent } from './feedback.component';

@NgModule({
  imports: [
    FeedbackRoutingModule,
    AppCommonModule,
  ],
  declarations: [FeedbackComponent]
})
export class FeedbackModule { }
