import { Component, OnInit, Inject } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { AnswerBoxService } from '../../../core/services/answer-box.service';
import { WordCloudService } from '../../../core/services/word-cloud.service';
import { QuestionBoxService } from '../../../core/services/question-box.service';
import { VotesBoxService } from '../../../core/services/votes-box.service';
import { ProcessingService } from '../../../core/services/processing.service';
import { NotificationsService } from '../../../core/services/notifications.service';

import { NewBoxModel } from '../../../core/models/new-box.model';
import { QuestionBoxModel } from '../../../core/models/question-box.model';
import { WordCloudModel } from '../../../core/models/word-cloud.model';
import { AnswerBoxModel } from '../../../core/models/answer-box.model';
import { VotesBoxModel } from '../../../core/models/votes-box.model';
import { UserModel } from '../../../core/models/user.model';
import { ColorModel } from '../../../core/models/color.model';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';
import { WebsocketMessageModel } from '../../../core/models/websocket-message.model';
import { BaseBoxModel } from '../../../core/models/base-box.model';

@Component({
  selector: 'app-confirm-box-creation-dialog',
  templateUrl: './confirm-box-creation-dialog.component.html',
  styleUrls: ['./confirm-box-creation-dialog.component.scss']
})
export class ConfirmBoxCreationDialogComponent implements OnInit {

  //import box to be created
  newBox: NewBoxModel;
  accountWebSocket: $WebSocket;

  constructor(
    private router: Router,
    private notificationsService: NotificationsService,
    private questionBoxService: QuestionBoxService,
    private answerBoxService: AnswerBoxService,
    private wordCloudService: WordCloudService,
    private votesBoxService: VotesBoxService,
    public processingService: ProcessingService,
    @Inject(MAT_DIALOG_DATA) public data: { newBox: NewBoxModel, accountWebSocket: $WebSocket },
    private dialogRef: MatDialogRef<ConfirmBoxCreationDialogComponent>
  ) {
    this.newBox = data.newBox;
    this.accountWebSocket = data.accountWebSocket;
  }

  private createVotesBox(votesBox: VotesBoxModel): Observable<any> {
    let response = new Observable<any>(observable => {
      this.votesBoxService.createVotesBox(votesBox)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarError("Could not create votes poll.")
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.broadcastNewBoxMessage(votesBox.BaseBox);
          this.router.navigate([`/dashboard/boxes/details/${votesBox.BaseBox.Code}`])
          observable.next(true);
        });
    });
    return response;
  }

  private createQuestionBox(questionBox: QuestionBoxModel): Observable<any> {
    let response = new Observable<any>(observable => {
      this.questionBoxService.createQuestionBox(questionBox)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarError("Could not create question poll.")
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.broadcastNewBoxMessage(questionBox.BaseBox);
          this.router.navigate([`/dashboard/boxes/details/${questionBox.BaseBox.Code}`])
          observable.next(true);
        });
    });
    return response;
  }

  private createAnswerBox(answerBox: AnswerBoxModel): Observable<any> {
    let response = new Observable<any>(observable => {
      this.answerBoxService.createAnswerBox(answerBox)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarError("Could not create answer poll.")
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.broadcastNewBoxMessage(answerBox.BaseBox);
          this.router.navigate([`/dashboard/boxes/details/${answerBox.BaseBox.Code}`]);
          observable.next(true);
        });
    });
    return response;
  }

  private createWordCloud(wordCloud: WordCloudModel): Observable<any> {
    let response = new Observable<any>(observable => {
      this.wordCloudService.createWordCloud(wordCloud)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarError("Could not create word cloud box.")
          return Observable.throw(error);
        })
        .subscribe(response => {
          this.broadcastNewBoxMessage(wordCloud.BaseBox);
          this.router.navigate([`/dashboard/boxes/details/${wordCloud.BaseBox.Code}`])
          observable.next(true);
        })
    });
    return response;
  }

  private notifyBoxCreateSuccess() {
    this.notificationsService.openSnackBarDefault("Box successfully created!");
  }

  createBox(newBox: NewBoxModel) {
    this.processingService.enableProcessingAnimation();
    switch (newBox.BaseBox.BoxTypeId) {
      case 1:
        let themeId = newBox.BaseBox.ThemeId;
        let wordCloudModel: any = {
          DefaultWord: newBox.DefaultWord,
          Header: newBox.Header,
          Theme: newBox.Theme,
          BaseBox: {
            ThemeId: themeId,
            BoxTypeId: 1,
            Code: newBox.BaseBox.Code,
            Theme: newBox.BaseBox.Theme,
          }
        }
        this.createWordCloud(wordCloudModel)
          .catch(error => {
            console.error("Could not create word cloud.", error);
            return Observable.throw(error);
          })
          .subscribe(response => {
            this.processingService.disableProcessingAnimation();
            this.closeDialog()
            this.notifyBoxCreateSuccess();
          });
        return;
      case 2:
        let questionBox: any = {
          Header: newBox.Header,
          Theme: newBox.Theme,
          BaseBox: {
            ThemeId: newBox.BaseBox.ThemeId,
            BoxTypeId: 2,
            Code: newBox.BaseBox.Code,
            Theme: newBox.BaseBox.Theme,
          }
        }
        this.createQuestionBox(questionBox)
          .catch(error => {
            console.error("Could not create question poll.", error);
            return Observable.throw(error);
          })
          .subscribe(response => {
            this.processingService.disableProcessingAnimation();
            this.closeDialog()
            this.notifyBoxCreateSuccess();
          });
        return;
      case 3:
        let answerBox: any = {
          BaseBox: {
            ThemeId: newBox.BaseBox.ThemeId,
            BoxTypeId: 3,
            Code: newBox.BaseBox.Code,
            Theme: newBox.BaseBox.Theme,
          },
          Questions: newBox.Questions
        }
        this.createAnswerBox(answerBox)
          .catch(error => {
            console.error("Could not create answer poll.", error);
            return Observable.throw(error);
          })
          .subscribe(response => {
            this.processingService.disableProcessingAnimation();
            this.closeDialog()
            this.notifyBoxCreateSuccess();
          });
        return;
      case 4:
        let votesBox: any = {
          BaseBox: {
            ThemeId: newBox.BaseBox.ThemeId,
            BoxTypeId: 4,
            Code: newBox.BaseBox.Code,
            Theme: newBox.BaseBox.Theme,
          },
          Questions: newBox.Questions
        }
        this.createVotesBox(votesBox)
          .catch(error => {
            console.error("Could not create votes poll.", error);
            return Observable.throw(error);
          })
          .subscribe(response => {
            this.processingService.disableProcessingAnimation();
            this.closeDialog()
            this.notifyBoxCreateSuccess();
          });
        return;
      default:
        this.notificationsService.openSnackBarDefault("Could not create poll. Invalid poll type.");
        return;
    }
  }

  broadcastNewBoxMessage(baseBox: BaseBoxModel) {
    let boxMessage: BaseBoxModel = { Code: baseBox.Code, IsLive: baseBox.IsLive, PhoneNumber: baseBox.PhoneNumber, BoxTypeId: baseBox.BoxTypeId };
    let boxString = JSON.stringify(boxMessage);
    let newBoxMessage: WebsocketMessageModel = { MessageType: "NewBox", MessageData: boxString }
    this.accountWebSocket.send(newBoxMessage);
  }

  notNow() {
    this.closeDialog();
  }

  closeDialog() {
    this.dialogRef.close();
  }

  ngOnInit() {
  }

}
