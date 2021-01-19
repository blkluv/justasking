import { Component, OnInit, Input } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { environment } from '../../../environments/environment';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { QuestionBoxService } from '../../core/services/question-box.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';

import { BaseBoxModel } from '../../core/models/base-box.model';
import { QuestionBoxModel } from '../../core/models/question-box.model';
import { QuestionBoxEntryVoteModel } from '../../core/models/question-box-entry-vote.model';
import { QuestionBoxEntryModel } from '../../core/models/question-box-entry.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

@Component({
  selector: 'app-question-box-details',
  templateUrl: './question-box-details.component.html',
  styleUrls: ['./question-box-details.component.scss']
})
export class QuestionBoxDetailsComponent implements OnInit {

  @Input('boxWebSocket') boxWebSocket: $WebSocket;
  @Input('accountWebSocket') accountWebSocket: $WebSocket;
  @Input('boxClientCount') boxClientCount: number;

  box: QuestionBoxModel;
  entries: any[]; 
  loadingEntries: boolean;
  loadingBox: boolean;
  ws: $WebSocket;

  constructor(
    private notificationsService: NotificationsService,
    private route: ActivatedRoute,
    private questionBoxService: QuestionBoxService,
    private webSocketService: WebSocketService,
    private processingService: ProcessingService
  ) {
    this.box = { BaseBox: {} };
    this.entries = [];
  }

  toggleEntry(entry: any) {
    this.toggleQuestionBoxEntryVisibility(entry);
  }

  hideAllEntries() {
    this.hideAllQuestionEntries();
  }

  unhideAllEntries() {
    this.unhideAllQuestionEntries();
  }

  private hideAllQuestionEntries() {
    let questionBox: QuestionBoxModel = {
      BoxId: this.box.BaseBox.ID
    }
    this.processingService.enableProcessingAnimation();
    this.questionBoxService.hideAllEntries(questionBox)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("Could not hide all entries.");
        return Observable.throw(error);
      })
      .subscribe(data => {
        this.processingService.disableProcessingAnimation();
        let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntryHideAll", MessageData: null }
        this.ws.send(message);
        this.entries.forEach(entry => {
          entry.IsHidden = true;
        });
        this.notificationsService.openSnackBarDefault("All entries are now hidden.");
      });
  }

  private unhideAllQuestionEntries() {
    let questionBox: QuestionBoxModel = {
      BoxId: this.box.BaseBox.ID
    }
    this.processingService.enableProcessingAnimation();
    this.questionBoxService.unhideAllEntries(questionBox)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("Could not unhide all entries.");
        return Observable.throw(error);
      })
      .subscribe(data => {
        this.processingService.disableProcessingAnimation();
        let entriesString = JSON.stringify(this.entries);
        let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntryUnhideAll", MessageData: entriesString }
        this.ws.send(message);
        this.entries.forEach(entry => {
          entry.IsHidden = false;
        });
        this.notificationsService.openSnackBarDefault("All entries are now displayed.");
      });
  }

  private toggleQuestionBoxEntryVisibility(entry: any) {
    this.processingService.enableProcessingAnimation();
    if (entry.IsHidden) {
      this.questionBoxService.unhideEntry(this.box.BaseBox.Code, entry)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          console.error("Could not unhide question poll entry visibility.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          entry.IsHidden = !entry.IsHidden;
          let entryString = JSON.stringify(entry);
          let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntryUnhide", MessageData: entryString }
          this.ws.send(message);
        });
    } else {
      this.questionBoxService.hideEntry(this.box.BaseBox.Code, entry)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          console.error("Could not hide question poll entry visibility.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          entry.IsHidden = !entry.IsHidden;
          let entryString = JSON.stringify(entry);
          let message: WebsocketMessageModel = { MessageType: "QuestionBoxEntryHide", MessageData: entryString }
          this.ws.send(message);
        });
    }
  }

  private initializeQuestionBoxDetails(code: string) {
    this.loadingEntries = true;
    this.questionBoxService.getEntries(code)
      .catch(error => {
        this.loadingEntries = false;
        return Observable.throw(error);
      })
      .subscribe(entries => {
        this.loadingEntries = false;
        this.entries = entries;
      });
    this.ws = this.boxWebSocket;
    this.ws.onMessage(this.onQuestionBoxHubMessage, { autoApply: false });
  }

  private onQuestionBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let entry = JSON.parse(response.MessageData);
    if (response.MessageType == "QuestionBoxEntry") {
      this.entries.push(entry)
    } else if (response.MessageType == "DashboardClientCount") {
      let count = JSON.parse(response.MessageData).ClientCount;
      this.boxClientCount = count;
    } else if (response.MessageType == "QuestionBoxEntryVote") {
      let vote: QuestionBoxEntryVoteModel = JSON.parse(response.MessageData);
      let affectedEntry: QuestionBoxEntryModel = this.entries.find(submittedEntry => submittedEntry.EntryId == vote.EntryId);
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
    } else if (response.MessageType == "QuestionBoxEntryHide") {
      let affectedEntry: QuestionBoxEntryModel = this.entries.find(submittedEntry => submittedEntry.EntryId == entry.EntryId);
      affectedEntry.IsHidden = true;
    } else if (response.MessageType == "QuestionBoxEntryUnhide") {
      let affectedEntry: QuestionBoxEntryModel = this.entries.find(submittedEntry => submittedEntry.EntryId == entry.EntryId);
      affectedEntry.IsHidden = false;
    }
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
      this.box.BaseBox.Code = code;
    });
    this.loadingBox = true;
    this.questionBoxService.getQuestionBox(code)
      .catch(error => {
        console.error("Could not get question poll to initialize box details.", error);
        return Observable.throw(error);
      })
      .subscribe(box => {
        this.loadingBox = false;
        this.box = box;
        this.initializeQuestionBoxDetails(code);
      });
  }
}

