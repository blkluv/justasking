import { NgModule } from '@angular/core';

import { AppCommonModule } from '../app-common/app-common.module';
import { AnswerBoxEntryComponent } from './answer-box-entry/answer-box-entry.component';
import { AnswerBoxPresentationComponent } from './answer-box-presentation/answer-box-presentation.component';
import { AnswerBoxDetailsComponent } from './answer-box-details/answer-box-details.component';

import { AnswerBoxRoutingModule } from './answer-box.routing.module';

@NgModule({
  imports: [
    AppCommonModule,
    AnswerBoxRoutingModule
  ],
  exports:[
    AnswerBoxEntryComponent, 
    AnswerBoxPresentationComponent,
    AnswerBoxDetailsComponent
  ],
  declarations: [
    AnswerBoxEntryComponent, 
    AnswerBoxPresentationComponent,
    AnswerBoxDetailsComponent
  ]
})
export class AnswerBoxModule { }
