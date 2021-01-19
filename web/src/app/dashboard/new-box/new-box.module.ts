import { NgModule } from '@angular/core';

import { AppCommonModule } from '../../app-common/app-common.module';
import { QuestionBoxModule } from '../../question-box/question-box.module';
import { AnswerBoxModule } from '../../answer-box/answer-box.module';
import { VotesBoxModule } from '../../votes-box/votes-box.module';
import { WordCloudModule } from '../../word-cloud/word-cloud.module';
import { NewBoxRoutingModule } from './new-box.routing.module';

import { BaseBoxModule } from '../../base-box/base-box.module';

import { NewBoxComponent } from './new-box.component';
import { ConfirmBoxCreationDialogComponent } from './confirm-box-creation-dialog/confirm-box-creation-dialog.component';
import { NewUserWelcomeDialogComponent } from './new-user-welcome-dialog/new-user-welcome-dialog.component';

@NgModule({
  imports: [
    NewBoxRoutingModule,
    AppCommonModule,
    BaseBoxModule,
    QuestionBoxModule,
    VotesBoxModule,
    AnswerBoxModule,
    WordCloudModule,
  ],
  declarations: [
    NewBoxComponent,
    NewUserWelcomeDialogComponent,
    ConfirmBoxCreationDialogComponent,
  ],
  entryComponents: [
    NewUserWelcomeDialogComponent ,
    ConfirmBoxCreationDialogComponent
  ]
})
export class NewBoxModule {}
