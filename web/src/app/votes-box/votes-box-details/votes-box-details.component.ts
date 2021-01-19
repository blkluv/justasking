import { Component, OnInit, Input } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { environment } from '../../../environments/environment';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { BoxService } from '../../core/services/box.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { VotesBoxService } from '../../core/services/votes-box.service';

import { BaseBoxModel } from '../../core/models/base-box.model';
import { VotesBoxModel } from '../../core/models/votes-box.model';
import { VotesBoxVoteModel } from '../../core/models/votes-box-vote.model';
import { VotesBoxQuestionModel } from '../../core/models/votes-box-question.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

@Component({
  selector: 'app-votes-box-details',
  templateUrl: './votes-box-details.component.html',
  styleUrls: ['./votes-box-details.component.scss']
})
export class VotesBoxDetailsComponent implements OnInit {

  @Input('boxWebSocket') boxWebSocket: $WebSocket;
  @Input('accountWebSocket') accountWebSocket: $WebSocket;
  @Input('boxClientCount') boxClientCount: number;

  box: VotesBoxModel; 
  votesCount: number;
  loadingEntries: boolean;
  loadingBox: boolean;
  ws: $WebSocket;
  answerIdToIndexAssossiativeArray: any; 

  constructor(
    private notificationsService: NotificationsService,
    private route: ActivatedRoute,
    private votesBoxService: VotesBoxService,
    private boxService: BoxService,
    private webSocketService: WebSocketService,
    private processingService: ProcessingService,
  ) {
    this.votesCount = 0;
    this.box = {};
    this.answerIdToIndexAssossiativeArray = {};
  }

  toggleVotesBoxQuestion(question: VotesBoxQuestionModel) {
    this.processingService.enableProcessingAnimation();
    if (question.IsActive) {
      this.votesBoxService.deActivateQuestion(question)
        .catch(error => {
          console.error("Could not deactivate question for answer poll.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          question.IsActive = !question.IsActive;
          let questionHubMessage: VotesBoxQuestionModel = {
            IsActive: question.IsActive,
            SortOrder: question.SortOrder,
            QuestionId: question.QuestionId
          };
          let questionString = JSON.stringify(questionHubMessage);
          let message: WebsocketMessageModel = { MessageType: "VotesBoxQuestionDeactivate", MessageData: questionString }
          this.ws.send(message);
          this.notificationsService.openSnackBarDefault("Question is no longer allowing votes.");
        });
    } else {
      this.votesBoxService.activateQuestion(question)
        .catch(error => {
          console.error("Could not activate question for answer poll.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          question.IsActive = !question.IsActive;
          let questionHubMessage: VotesBoxQuestionModel = {
            Answers: question.Answers,
            BoxId: question.BoxId, 
            Header: question.Header,
            IsActive: question.IsActive,
            QuestionId: question.QuestionId,
            SortOrder: question.SortOrder,
            SelectedAnswerId: null
          };
          let questionString = JSON.stringify(questionHubMessage);
          let message: WebsocketMessageModel = { MessageType: "VotesBoxQuestionActivate", MessageData: questionString }
          this.ws.send(message);
          this.notificationsService.openSnackBarDefault("Question is allowing votes.");
        });
    }
  }

  private onVotesBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let votes: VotesBoxVoteModel[] = JSON.parse(response.MessageData);
    if (response.MessageType == "VotesBoxEntrySubmit") {
      this.addVoteToVotesBoxGraph(votes);
    } else if (response.MessageType == "DashboardClientCount") {
      let count = JSON.parse(response.MessageData).ClientCount;
      this.boxClientCount = count;
    } else if (response.MessageType == "VotesBoxQuestionActivate"){
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = broadcastedQuestion.IsActive;
        }
      });
    } else if (response.MessageType == "VotesBoxQuestionDeactivate"){
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = broadcastedQuestion.IsActive;
        }
      });
    }
  }

  private addVoteToVotesBoxGraph(votes) {
    votes.forEach((vote: VotesBoxVoteModel) => {
      this.box.Questions.forEach(question => {
        question.Answers.forEach(answer => {
          let graphIndex = this.answerIdToIndexAssossiativeArray[answer.AnswerId];
          if ((vote.QuestionId == question.QuestionId) && (vote.AnswerId == answer.AnswerId)) {
            answer.Votes++;
            this.votesCount++; 
          }
        });
        question.Answers = question.Answers.slice();
      });
    });
  }

  ngOnInit() { 
    this.loadingEntries = true;
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });
    this.votesBoxService.getVotesBox(code)
      .catch(error => {
        this.loadingEntries = false;
        return Observable.throw(error);
      })
      .subscribe((votesBox: VotesBoxModel) => {
        this.loadingEntries = false;
        this.box = votesBox;
        this.box.Questions.forEach(question => {
          question.Answers.forEach(answer => {
            this.votesCount += answer.Votes;
          });
        });
      });
    this.ws = this.boxWebSocket;
    this.ws.onMessage(this.onVotesBoxHubMessage, { autoApply: false });
  }
}