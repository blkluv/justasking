import { Component, Input, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';
import { UUID } from 'angular2-uuid';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { NotificationsService } from '../../core/services/notifications.service';
import { VotesBoxService } from '../../core/services/votes-box.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { RandomizerService } from '../../core/services/randomizer.service';

import { VotesBoxModel } from '../../core/models/votes-box.model';
import { VotesBoxVoteModel } from '../../core/models/votes-box-vote.model';
import { VotesBoxQuestionModel } from '../../core/models/votes-box-question.model';
import { VotesBoxQuestionAnswerModel } from '../../core/models/votes-box-question-answer.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-votes-box-presentation',
  templateUrl: './votes-box-presentation.component.html',
  styleUrls: ['./votes-box-presentation.component.scss']
})
export class VotesBoxPresentationComponent implements OnInit {

  @Input('box') baseBox: BaseBoxModel;
  @Input('model') votesBoxModel: VotesBoxModel;
  @Input('boxWebSocket') boxWebSocket: $WebSocket;

  ws: $WebSocket;
  box: VotesBoxModel;
  visibleQuestionsCount: number = 0;
  answerIdToIndexAssossiativeArray: any;
  loading: boolean;

  constructor(
    private notificationsService: NotificationsService,
    private votesBoxService: VotesBoxService,
    private randomizerService: RandomizerService,
    private router: Router,
    private route: ActivatedRoute,
    private webSocketService: WebSocketService
  ) {
    this.loading = false;
    this.box = { BaseBox: {} };
    this.answerIdToIndexAssossiativeArray = {};
  }

  private onVotesBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let votes: VotesBoxVoteModel[] = JSON.parse(response.MessageData);
    if (response.MessageType == "VotesBoxEntrySubmit") {
      this.addVoteToVotesBoxGraph(votes);
    }
    if (response.MessageType == "VotesBoxQuestionDeactivate") {
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if (question.QuestionId == broadcastedQuestion.QuestionId) {
          question.IsActive = broadcastedQuestion.IsActive;
          this.syncQuestionsCount()
        }
      });
    } else if (response.MessageType == "VotesBoxQuestionActivate") {
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if (question.QuestionId == broadcastedQuestion.QuestionId) {
          question.IsActive = broadcastedQuestion.IsActive;
          this.syncQuestionsCount();
        }
      });
    }
  }

  private addVoteToVotesBoxGraph(votes) {
    votes.forEach((vote: VotesBoxVoteModel) => {
      if (this.box.Questions) {
        this.box.Questions.forEach(question => {
          question.Answers.forEach(answer => {
            let graphIndex = this.answerIdToIndexAssossiativeArray[answer.AnswerId];
            if ((vote.QuestionId == question.QuestionId) && (vote.AnswerId == answer.AnswerId)) {
              answer.Votes++;
            }
          });
          question.Answers = question.Answers.slice();
        });
      }
    });
  }

  private syncQuestionsCount() {
    this.visibleQuestionsCount = 0;
    this.box.Questions.forEach(question => {
      question.SelectedAnswerId = null;
      if (question.IsActive) {
        this.visibleQuestionsCount++;
      }
    });
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });

    if (!this.votesBoxModel.IsPreview) {
      this.box.BaseBox = this.votesBoxModel.BaseBox;
      this.loading = true;
      this.votesBoxService.getVotesBox(code)
        .catch(error => {
          this.loading = false;
          console.error("Could not get votes poll.", error);
          return Observable.throw(error);
        })
        .subscribe((votesBox: VotesBoxModel) => {
          this.loading = false;
          this.box = votesBox;
          this.box.Questions.forEach(question => {
            if (question.IsActive) {
              this.visibleQuestionsCount++;
            }
          });
        });
      this.ws = this.boxWebSocket;
      this.ws.onMessage(this.onVotesBoxHubMessage, { autoApply: false });
    } else {
      //preview screen
      this.box = this.votesBoxModel;
      //populate answers count
      this.box.Questions.forEach(question => {
        question.Answers.forEach(answer => {
          answer.Votes = this.randomizerService.getRandomVotesBoxEntries(50);
        });
      });
      //initializing preview graphs
      this.box.Questions.forEach(question => {
        if (question.IsActive) {
          this.visibleQuestionsCount++;
        }
      });
    }
  }
}
