import { NgModule } from '@angular/core';
import { AppCommonModule } from '../app-common/app-common.module';
import { EntryBoxComponent } from './entry-box/entry-box.component';
import { PresentationBoxComponent } from './presentation-box/presentation-box.component';

import { BaseBoxRoutingModule } from './base-box.routing.module';
import { QuestionBoxModule } from '../question-box/question-box.module';
import { VotesBoxModule } from '../votes-box/votes-box.module';
import { AnswerBoxModule } from '../answer-box/answer-box.module';
import { WordCloudModule } from '../word-cloud/word-cloud.module';

@NgModule({
  imports: [
    AppCommonModule,
    BaseBoxRoutingModule,
    QuestionBoxModule,
    VotesBoxModule,
    AnswerBoxModule,
    WordCloudModule
  ],
  exports:[
    EntryBoxComponent, 
    PresentationBoxComponent
  ],
  declarations: [
    EntryBoxComponent, 
    PresentationBoxComponent
  ]
})
export class BaseBoxModule { }
