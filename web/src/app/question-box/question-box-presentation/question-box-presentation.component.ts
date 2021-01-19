import { Component, Input, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { environment } from '../../../environments/environment';

import { RandomizerService } from '../../core/services/randomizer.service';
import { QuestionBoxService } from '../../core/services/question-box.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { RelevancePipe } from '../../core/pipes/relevance.pipe';

import { QuestionBoxEntryVoteModel } from '../../core/models/question-box-entry-vote.model';
import { QuestionBoxModel } from '../../core/models/question-box.model';
import { QuestionBoxEntryModel } from '../../core/models/question-box-entry.model';
import { ColorModel } from '../../core/models/color.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-question-box-presentation',
  templateUrl: './question-box-presentation.component.html',
  styleUrls: ['./question-box-presentation.component.scss']
})
export class QuestionBoxPresentationComponent implements OnInit {

  @Input('box') baseBox: BaseBoxModel;
  @Input('model') questionBoxModel: QuestionBoxModel;
  @Input('boxWebSocket') boxWebSocket: $WebSocket;

  box: QuestionBoxModel;
  submitted: QuestionBoxEntryModel[];
  response: any;
  loading: boolean;

  constructor(
    private questionBoxService: QuestionBoxService,
    private randomizerService: RandomizerService,
    private relevancePipe: RelevancePipe,
    private router: Router,
    private route: ActivatedRoute,
    private webSocketService: WebSocketService
  ) {
    this.loading = false;
    this.submitted = [];
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

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });

    if (!this.questionBoxModel.IsPreview) {
      this.box = { BaseBox: this.questionBoxModel.BaseBox };
      this.loading = true;
      this.questionBoxService.getQuestionBox(code)
        .catch(error => {
          this.loading = false;
          console.error("Could not get question poll.", error);
          return Observable.throw(error);
        })
        .subscribe(questionbox => {
          this.loading = false;
          this.questionBoxModel = questionbox;
          this.box = questionbox;
        });
      this.questionBoxService.getVisibleEntries(code)
        .catch(error => {
          console.error("Could not get question poll entries.", error);
          return Observable.throw(error);
        })
        .subscribe(entries => {
          this.submitted = this.relevancePipe.transform(entries);
        });
      let ws = this.boxWebSocket;
      ws.onMessage(this.onQuestionBoxHubMessage, { autoApply: false });
    } else {
      this.box = this.questionBoxModel;
      this.submitted = this.randomizerService.getRandomQuestionBoxEntries(4);
    }
  }
}
