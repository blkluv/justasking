import { Component, Input, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { environment } from '../../../environments/environment';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ColorModel } from '../../core/models/color.model';
import { QuestionBoxService } from '../../core/services/question-box.service';
import { RelevancePipe } from '../../core/pipes/relevance.pipe';
import { QuestionBoxModel } from '../../core/models/question-box.model';
import { QuestionBoxEntryModel } from '../../core/models/question-box-entry.model';
import { QuestionBoxEntryVoteModel } from '../../core/models/question-box-entry-vote.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { WebSocketService } from '../../core/services/web-socket.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-question-box-entry',
  templateUrl: './question-box-entry.component.html',
  styleUrls: ['./question-box-entry.component.scss']
})
export class QuestionBoxEntryComponent implements OnInit {

  @Input('box') baseBox: BaseBoxModel;
  @Input('model') questionBoxModel: QuestionBoxModel;
  @Input('boxWebSocket') boxWebSocket: $WebSocket;

  submitQuestionForm: any;
  box: QuestionBoxModel;
  submitted: QuestionBoxEntryModel[];
  questionBoxEntryModel: QuestionBoxEntryModel;
  ws: $WebSocket;
  loading: boolean;

  constructor(
    private notificationsService: NotificationsService,
    private questionBoxService: QuestionBoxService,
    private relevancePipe: RelevancePipe,
    private router: Router,
    private route: ActivatedRoute,
    private webSocketService: WebSocketService
  ) {
    this.loading = false;
    this.questionBoxEntryModel = { Question: "" }
    this.box = { Question: "", BaseBox: {} };
    this.submitted = [];
  }

  onSubmit() {
    let entry: QuestionBoxEntryModel = {
      BoxId: this.questionBoxEntryModel.BoxId,
      Code: this.questionBoxEntryModel.Code,
      Question: this.questionBoxEntryModel.Question,
      Upvotes: 1,
      Downvotes: 0
    }
    this.questionBoxService.submitEntry(entry)
      .catch(error => {
        console.error("Could not submit entry for question poll.", error);
        return Observable.throw(error);
      })
      .subscribe(response => {
        let entryString = JSON.stringify(response);
        let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntry", MessageData: entryString }
        this.ws.send(message);
        this.box.Question = "";
        this.notificationsService.openSnackBarDefault("Thank you for your question!");
        let entryEl = document.getElementById("boxQuestion");
        entryEl.style.height = "42px";
        entryEl.focus();
        this.displayUserVoteActions(this.box.BaseBox.Code);
        //this.submitted = this.relevancePipe.transform(this.submitted);
      });
  }

  upvote(entry: QuestionBoxEntryModel) {
    let vote: QuestionBoxEntryVoteModel = new QuestionBoxEntryVoteModel;
    vote.EntryId = entry.EntryId;
    vote.VoteType = "upvote"
    vote.VoteValue = 1;
    vote.Code = this.questionBoxEntryModel.Code;
    this.questionBoxService.vote(vote)
      .catch(error => {
        console.error("Could not upvote for entry for question poll.", error);
        return Observable.throw(error);
      })
      .subscribe(response => {
        entry.UserUpvoted = false;
        entry.UserDownvoted = true;
        let voteString = JSON.stringify(response);
        let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntryVote", MessageData: voteString }
        this.ws.send(message);
        this.displayUserVoteActions(this.box.BaseBox.Code);
        this.submitted = this.relevancePipe.transform(this.submitted);
      });
  }

  downvote(entry: QuestionBoxEntryModel) {
    let vote: QuestionBoxEntryVoteModel = new QuestionBoxEntryVoteModel;
    vote.EntryId = entry.EntryId;
    vote.VoteType = "downvote"
    vote.VoteValue = 1;
    vote.Code = this.questionBoxEntryModel.Code;
    this.questionBoxService.vote(vote)
      .catch(error => {
        console.error("Could not downvote for entry for question poll.", error);
        return Observable.throw(error);
      })
      .subscribe(response => {
        entry.UserUpvoted = true;
        entry.UserDownvoted = false;
        let voteString = JSON.stringify(response);
        let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntryVote", MessageData: voteString }
        this.ws.send(message);
        this.displayUserVoteActions(this.box.BaseBox.Code);
        this.submitted = this.relevancePipe.transform(this.submitted);
      });
  }

  private removeFromQuestionBoxEntries(entry: QuestionBoxEntryModel) {
    this.submitted.forEach((submittedEntry, index) => {
      if (submittedEntry.EntryId == entry.EntryId) {
        this.submitted.splice(index, 1);
      }
    })
  }

  private addToQuestionBoxEntries(entry: QuestionBoxEntryModel) {
    this.submitted.push(entry);
    this.submitted = this.relevancePipe.transform(this.submitted);
  }

  private onQuestionBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    if (response.MessageType == "QuestionBoxEntry") {
      let entry = JSON.parse(response.MessageData);
      this.submitted.push(entry);
      this.submitted = this.relevancePipe.transform(this.submitted);
    } else if (response.MessageType == "QuestionBoxEntryVote") {
      let vote: QuestionBoxEntryVoteModel = JSON.parse(response.MessageData);
      let affectedEntry: QuestionBoxEntryModel = this.submitted.find(submittedEntry => submittedEntry.EntryId == vote.EntryId);
      if (vote.IsUndo) {
        if (vote.VoteType == "upvote") {
          affectedEntry.Upvotes -= 1;
        }
        else if (vote.VoteType == "downvote") {
          affectedEntry.Downvotes -= 1;
        }
      } else if (vote.IsDownvoteFromUpvote) {
        affectedEntry.Upvotes -= 1;
        affectedEntry.Downvotes += 1;
      } else if (vote.IsUpvoteFromDownvote) {
        affectedEntry.Downvotes -= 1;
        affectedEntry.Upvotes += 1;
      } else {
        if (vote.VoteType == "upvote") {
          affectedEntry.Upvotes += 1;
        }
        else if (vote.VoteType == "downvote") {
          affectedEntry.Downvotes += 1;
        }
      }
      this.submitted = this.relevancePipe.transform(this.submitted);
    } else if (response.MessageType == "QuestionBoxEntryHide") {
      let entry = JSON.parse(response.MessageData);
      this.removeFromQuestionBoxEntries(entry);
    } else if (response.MessageType == "QuestionBoxEntryUnhide") {
      let entry = JSON.parse(response.MessageData);
      this.addToQuestionBoxEntries(entry);
    } else if (response.MessageType == "QuestionBoxEntryHideAll") {
      this.submitted = [];
    } else if (response.MessageType == "QuestionBoxEntryUnhideAll") {
      let entries = JSON.parse(response.MessageData);
      this.submitted = entries;
      this.submitted = this.relevancePipe.transform(this.submitted);
    } else if (response.MessageType == "OpenBox") {
      let box = JSON.parse(response.MessageData);
      this.box.BaseBox.PhoneNumber = box.PhoneNumber;
    } else if (response.MessageType == "CloseBox") {
      let box = JSON.parse(response.MessageData);
      this.box.BaseBox.PhoneNumber = box && box.PhoneNumber;
    }
  }

  private displayUserVoteActions(code: string) {
    this.questionBoxService.getAllUserVotesFromLocal(code)
      .catch(error => {
        console.error("Could not get all user votes from browser repo for question poll.", error);
        return Observable.throw(error);
      })
      .subscribe(userVotes => {
        let associateVotes = {};
        userVotes.forEach((vote: QuestionBoxEntryVoteModel) => {
          associateVotes[vote.EntryId] = vote;
        });
        this.submitted.forEach(entry => {
          entry.UserUpvoted = false;
          entry.UserDownvoted = false;
          if (associateVotes && associateVotes[entry.EntryId]) {
            if (associateVotes[entry.EntryId].VoteType == 'upvote') {
              entry.UserUpvoted = true;
              entry.UserDownvoted = false;
            } else if (associateVotes[entry.EntryId].VoteType == 'downvote') {
              entry.UserDownvoted = true;
              entry.UserUpvoted = false;
            }
          }
        });
      });
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });
    if (!this.questionBoxModel.IsPreview) {
      this.box.BaseBox = this.baseBox;
      this.loading = true;
      this.questionBoxService.getQuestionBox(code)
        .catch(error => {
          this.loading = false;
          console.error("Could not get question poll.", error);
          return Observable.throw(error);
        })
        .subscribe(questionBox => {
          this.loading = false;
          this.questionBoxEntryModel = questionBox;
          this.questionBoxEntryModel.BoxId = questionBox.BoxId;
          this.questionBoxEntryModel.Code = questionBox.BaseBox.Code;
          this.box = questionBox;
          this.box.Question = "";
        });
      this.questionBoxService.getVisibleEntries(code)
        .catch(error => {
          console.error("Could not get question poll entries.", error);
          return Observable.throw(error);
        })
        .subscribe(entries => {
          this.submitted = entries;
          this.displayUserVoteActions(code);
          this.submitted = this.relevancePipe.transform(this.submitted);
        });
      this.ws = this.boxWebSocket;
      this.ws.onMessage(this.onQuestionBoxHubMessage, { autoApply: false }
      );
    } else {
      this.box = this.questionBoxModel;
      this.box.Question = "";
    }
  }
}
