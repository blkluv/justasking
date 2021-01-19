import { NgModule } from '@angular/core';

import { AppCommonModule } from '../app-common/app-common.module';

import { QuestionBoxRoutingModule } from './question-box.routing.module';
import { QuestionBoxEntryComponent } from './question-box-entry/question-box-entry.component';
import { QuestionBoxPresentationComponent } from './question-box-presentation/question-box-presentation.component';
import { QuestionBoxDetailsComponent } from './question-box-details/question-box-details.component';

@NgModule({
  imports: [
    QuestionBoxRoutingModule,
    AppCommonModule,
  ],
  exports:[
    QuestionBoxEntryComponent,
    QuestionBoxPresentationComponent,
    QuestionBoxDetailsComponent
  ],
  declarations: [
    QuestionBoxEntryComponent,
    QuestionBoxPresentationComponent,
    QuestionBoxDetailsComponent
  ]
})
export class QuestionBoxModule { }
