import { Component, OnInit, Input } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { environment } from '../../../environments/environment';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { AnswerBoxService } from '../../core/services/answer-box.service';
import { BoxService } from '../../core/services/box.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';

import { QuestionModel } from '../../core/models/question.model';
import { BaseBoxModel } from '../../core/models/base-box.model';
import { AnswerBoxModel } from '../../core/models/answer-box.model';
import { VotesBoxVoteModel } from '../../core/models/votes-box-vote.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

@Component({
  selector: 'app-answer-box-details',
  templateUrl: './answer-box-details.component.html',
  styleUrls: ['./answer-box-details.component.scss']
})
export class AnswerBoxDetailsComponent implements OnInit {

  @Input('boxWebSocket') boxWebSocket: $WebSocket;
  @Input('accountWebSocket') accountWebSocket: $WebSocket;
  @Input('boxClientCount') boxClientCount: number;

  box: AnswerBoxModel;
  sectionIndexes: any;
  entries: any[];
  loadingEntries: boolean;
  loadingBox: boolean;
  ws: $WebSocket;

  constructor(
    private notificationsService: NotificationsService,
    private route: ActivatedRoute,
    private answerBoxService: AnswerBoxService,
    private boxService: BoxService,
    private webSocketService: WebSocketService,
    private processingService: ProcessingService
  ) {
    this.box = {};
    this.entries = [];
  }

  toggleEntry(entry: any) {
    this.toggleAnswerBoxEntryVisibility(entry);
  }

  hideAllAnswerBoxEntries(question: QuestionModel) {
    this.processingService.enableProcessingAnimation();
    this.answerBoxService.hideAllEntries(question)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("Could not hide all entries for this question.");
        return Observable.throw(error);
      })
      .subscribe(data => {
        this.processingService.disableProcessingAnimation();
        let questionHubMessage: QuestionModel = {
          QuestionId: question.QuestionId,
          SortOrder: question.SortOrder,
        }
        let questionString = JSON.stringify(questionHubMessage);
        let message: WebsocketMessageModel = { MessageType: "AnswerBoxQuestionHideAll", MessageData: questionString }
        this.ws.send(message);
        question.Entries.forEach(entry => {
          entry.IsHidden = true;
        });
        this.notificationsService.openSnackBarDefault("All entries for this question are now hidden.");
      });
  }

  unhideAllAnswerBoxEntries(question: QuestionModel) {
    this.processingService.enableProcessingAnimation();
    this.answerBoxService.unhideAllEntries(question)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("Could not unhide all entries for this question.");
        return Observable.throw(error);
      })
      .subscribe(data => {
        this.processingService.disableProcessingAnimation();
        let questionHubMessage: QuestionModel = {
          QuestionId: question.QuestionId,
          SortOrder: question.SortOrder,
          Entries: question.Entries
        }
        let questionString = JSON.stringify(questionHubMessage);
        let message: WebsocketMessageModel = { MessageType: "AnswerBoxQuestionUnhideAll", MessageData: questionString }
        this.ws.send(message);
        question.Entries.forEach(entry => {
          entry.IsHidden = false;
        });
        this.notificationsService.openSnackBarDefault("All entries for this question are now displayed.");
      });
  }

  toggleAnswerBoxQuestion(question: QuestionModel) {
    this.processingService.enableProcessingAnimation();
    if (question.IsActive) {
      this.answerBoxService.deActivateQuestion(question)
        .catch(error => {
          console.error("Could not deactivate question for answer poll.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          question.IsActive = !question.IsActive;
          let questionHubMessage: QuestionModel = {
            Entry: null,
            Entries: null,
            IsActive: question.IsActive,
            Question: question.Question,
            SortOrder: question.SortOrder,
            QuestionId: question.QuestionId,
            Collapsed: question.Collapsed
          };
          let questionString = JSON.stringify(questionHubMessage);
          let message: WebsocketMessageModel = { MessageType: "AnswerBoxQuestionDeactivate", MessageData: questionString }
          this.ws.send(message);
          this.notificationsService.openSnackBarDefault("Question is no longer allowing entries.");
        });
    } else {
      this.answerBoxService.activateQuestion(question)
        .catch(error => {
          console.error("Could not activate question for answer poll.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          question.IsActive = !question.IsActive;
          let questionHubMessage: QuestionModel = {
            Entry: null,
            Entries: null,
            IsActive: question.IsActive,
            Question: question.Question,
            SortOrder: question.SortOrder,
            QuestionId: question.QuestionId,
            Collapsed: question.Collapsed
          };
          let questionString = JSON.stringify(questionHubMessage);
          let message: WebsocketMessageModel = { MessageType: "AnswerBoxQuestionActivate", MessageData: questionString }
          this.ws.send(message);
          this.notificationsService.openSnackBarDefault("Question is allowing entries.");
        });
    }
  }

  private toggleAnswerBoxEntryVisibility(entry: any) {
    this.processingService.enableProcessingAnimation();
    if (entry.IsHidden) {
      this.answerBoxService.unhideEntry(entry)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          console.error("Could not unhide answer poll entry visibility.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          entry.IsHidden = !entry.IsHidden;
          let entryString = JSON.stringify(entry);
          let message: WebsocketMessageModel = { MessageType: "AnswerBoxEntryUnhide", MessageData: entryString }
          this.ws.send(message);
        });
    } else {
      this.answerBoxService.hideEntry(entry)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          console.error("Could not hide answer poll entry visibility.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          entry.IsHidden = !entry.IsHidden;
          let entryString = JSON.stringify(entry);
          let message: WebsocketMessageModel = { MessageType: "AnswerBoxEntryHide", MessageData: entryString }
          this.ws.send(message);
        });
    }
  }

  private initializeAnswerBoxDetails(code: string) {
    this.loadingEntries = true;
    this.loadingBox = true;
    this.answerBoxService.getAnswerBox(code)
      .catch(error => {
        this.loadingBox = false;
        this.notificationsService.openSnackBarDefault("Could not get answer poll to initialize box details.");
        return Observable.throw(error);
      })
      .subscribe(answerBox => {
        this.loadingBox = false;
        this.box = answerBox;
        this.answerBoxService.getAnswerBoxEntries(code)
          .catch(error => {
            this.loadingEntries = false;
            return Observable.throw(error);
          })
          .subscribe(entries => {
            this.loadingEntries = false;
            this.entries = entries;
            this.sectionIndexes = this.answerBoxService.getQuestionIndexesAssociativeArray(this.box.Questions);
            this.box.Questions = this.answerBoxService.mapQuestionsWithEntries(this.sectionIndexes, this.box.Questions, entries);
          });
      })
    this.ws = this.boxWebSocket;
    this.ws.onMessage(this.onAnswerBoxHubMessage, { autoApply: false });
  }

  private onAnswerBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let entries = JSON.parse(response.MessageData);
    if (response.MessageType == "AnswerBoxEntry") {
      //updates count
      this.entries = this.entries.concat(entries);

      //updates displayed entries
      entries.forEach(entry => {
        let questionIndex = this.sectionIndexes[entry.QuestionId];
        if (typeof (this.box.Questions[questionIndex].Entries) == 'undefined') {
          this.box.Questions[questionIndex].Entries = [entry];
        } else {
          this.box.Questions[questionIndex].Entries.push(entry);
        }
      });
    } else if (response.MessageType == "DashboardClientCount") {
      let count = JSON.parse(response.MessageData).ClientCount;
      this.boxClientCount = count;
    } else if (response.MessageType == "AnswerBoxEntryHide") {
      let entry = JSON.parse(response.MessageData);
      let questionIndex = this.sectionIndexes[entry.QuestionId];
      this.box.Questions[questionIndex].Entries.forEach((submittedEntry,index) => {
        if(submittedEntry.EntryId == entry.EntryId){
          this.box.Questions[questionIndex].Entries[index].IsHidden = true;
        }
      });
    } else if (response.MessageType == "AnswerBoxEntryUnhide") {
      let entry = JSON.parse(response.MessageData);
      let questionIndex = this.sectionIndexes[entry.QuestionId];
      this.box.Questions[questionIndex].Entries.forEach((submittedEntry,index) => {
        if(submittedEntry.EntryId == entry.EntryId){
          this.box.Questions[questionIndex].Entries[index].IsHidden = false;
        }
      });
    } else if (response.MessageType == "AnswerBoxQuestionHideAll") {
      let entry = JSON.parse(response.MessageData);
      let questionIndex = this.sectionIndexes[entry.QuestionId];
      this.box.Questions[questionIndex].Entries.forEach(entry => {
        entry.IsHidden = true;
      });
    } else if (response.MessageType == "AnswerBoxQuestionUnhideAll") {
      let entry = JSON.parse(response.MessageData);
      let questionIndex = this.sectionIndexes[entry.QuestionId];
      this.box.Questions[questionIndex].Entries.forEach(entry => {
        entry.IsHidden = false;
      });
    } else if (response.MessageType == "AnswerBoxQuestionActivate") {
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = true;
        }
      });
    } else if (response.MessageType == "AnswerBoxQuestionDeactivate") {
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = false;
        }
      });
    }
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });
    this.initializeAnswerBoxDetails(code);
  }
}
