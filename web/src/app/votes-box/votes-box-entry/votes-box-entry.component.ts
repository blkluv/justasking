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
import { VotesBoxVoteRepoService } from '../../core/repos/votes-box-vote-repo.service';
import { WebSocketService } from '../../core/services/web-socket.service';

import { VotesBoxModel } from '../../core/models/votes-box.model';
import { VotesBoxVoteModel } from '../../core/models/votes-box-vote.model';
import { VotesBoxQuestionModel } from '../../core/models/votes-box-question.model';
import { VotesBoxQuestionAnswerModel } from '../../core/models/votes-box-question-answer.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-votes-box-entry',
  templateUrl: './votes-box-entry.component.html',
  styleUrls: ['./votes-box-entry.component.scss']
})
export class VotesBoxEntryComponent implements OnInit {

  @Input('box') baseBox: BaseBoxModel;
  @Input('model') votesBoxModel: VotesBoxModel;
  @Input('boxWebSocket') boxWebSocket: $WebSocket;

  allQuestionsVotedOn: boolean;
  userAlreadyVoted: boolean;
  box: VotesBoxModel;
  ws: $WebSocket;
  loading: boolean;

  constructor(
    private notificationsService: NotificationsService,
    private votesBoxService: VotesBoxService,
    private router: Router,
    private route: ActivatedRoute,
    private votesBoxVoteRepoService: VotesBoxVoteRepoService,
    private webSocketService: WebSocketService
  ) {
    this.loading = false;
    this.box = { BaseBox: {} };
    this.userAlreadyVoted = false;
    this.allQuestionsVotedOn = false;
  }

  onSubmit() {
    let votes: VotesBoxVoteModel[] = [];
    this.box.Questions.forEach(question => {
      votes.push({ QuestionId: question.QuestionId, AnswerId: question.SelectedAnswerId, Code: this.box.BaseBox.Code });
    });
    this.votesBoxService.submitVotes(votes)
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe((votesReponse: VotesBoxVoteModel[]) => {
        this.userAlreadyVoted = true;
        let votesString = JSON.stringify(votesReponse);
        let message: WebsocketMessageModel = { MessageType: "VotesBoxEntrySubmit", MessageData: votesString }
        this.ws.send(message);
        new Observable<any>(observable => {
          votes.forEach(vote => { this.votesBoxVoteRepoService.insert(vote).subscribe() })
          observable.next();
        })
        .catch(error=>{
          this.goToResults();
          return Observable.throw(error);
        })
        .subscribe(()=>{
          this.goToResults();
        });
      });
  }

  private goToResults(){
    this.webSocketService.disconnectWebSockets();
    // this.router.navigate([`${this.box.BaseBox.Code}/presentation`]);
  }

  selectAnswerForQuestion(question: VotesBoxQuestionModel, answer: VotesBoxQuestionAnswerModel) {
    question.SelectedAnswerId = answer.AnswerId;
    this.allQuestionsVotedOn = this.box.Questions.every(this.questionHasAnswerSelected);
  }

  private getSubmittedVotesAssosiativeArray(submittedVotes: VotesBoxVoteModel[]): any {
    let submittedVotesAssosiativeArray = {}
    submittedVotes.forEach(submittedVote => {
      submittedVotesAssosiativeArray[submittedVote.QuestionId] = true;
    });
    return submittedVotesAssosiativeArray;
  }

  private questionHasAnswerSelected(question: VotesBoxQuestionModel, index, array): boolean {
    let questionHasAnswerSelected = (question.IsActive == false) || (question.SelectedAnswerId && true);
    return questionHasAnswerSelected;
  }

  private resetForm() {
    this.userAlreadyVoted = false;
    this.box.Questions.forEach(question => {
      question.SelectedAnswerId = null;
    });
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    }); 
    if (!this.votesBoxModel.IsPreview) {
      this.ws = this.boxWebSocket;
      this.box.BaseBox = this.votesBoxModel.BaseBox;
      this.loading = true;
      // get votes poll details from the api
      this.votesBoxService.getVotesBox(code)
        .catch(error => {
          this.loading = false;
          console.error("Could not get votes poll.", error);
          return Observable.throw(error);
        })
        .subscribe(votesBox => {
          this.box = votesBox;
          //this.loading = false;          
          this.votesBoxVoteRepoService.getAllByCode(this.box.BaseBox.Code)
            .catch(error => {
              this.loading = false;
              console.error("Could not get submitted votes for this poll by this user.", error);
              return Observable.throw(error);
            })
            .subscribe((submittedVotes: VotesBoxVoteModel[]) => {
              this.loading = false;
              let submittedVotesAssosiativeArray = this.getSubmittedVotesAssosiativeArray(submittedVotes);
              this.box.Questions.forEach(question => {
                if (submittedVotesAssosiativeArray[question.QuestionId]) {
                  this.userAlreadyVoted = true;
                }
              });
            });
        });
    } else {
      this.box = this.votesBoxModel;
    }
  }
}