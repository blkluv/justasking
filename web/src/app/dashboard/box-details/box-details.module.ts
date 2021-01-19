import { NgModule } from '@angular/core';
 
import { AppCommonModule } from '../../app-common/app-common.module';

import { QuestionBoxModule } from '../../question-box/question-box.module';
import { AnswerBoxModule } from '../../answer-box/answer-box.module';
import { VotesBoxModule } from '../../votes-box/votes-box.module';
import { WordCloudModule } from '../../word-cloud/word-cloud.module';

import { BoxDetailsComponent } from './box-details.component';
import { BoxDetailsRoutingModule } from './box-details.routing.module'; 

@NgModule({
  imports: [
    BoxDetailsRoutingModule,
    AppCommonModule, 
    QuestionBoxModule,
    VotesBoxModule,
    AnswerBoxModule,
    WordCloudModule
  ],
  declarations: [
    BoxDetailsComponent
  ],
})
export class BoxDetailsModule { }
