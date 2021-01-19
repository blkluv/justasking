import { Component, OnInit, Input } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { environment } from '../../../environments/environment';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { UniqueEntriesPipe } from '../../core/pipes/unique-entries.pipe';

import { WordCloudService } from '../../core/services/word-cloud.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';

import { BaseBoxModel } from '../../core/models/base-box.model';
import { WordCloudModel } from '../../core/models/word-cloud.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

@Component({
  selector: 'app-word-cloud-details',
  templateUrl: './word-cloud-details.component.html',
  styleUrls: ['./word-cloud-details.component.scss']
})
export class WordCloudDetailsComponent implements OnInit {

  @Input('boxWebSocket') boxWebSocket: $WebSocket;
  @Input('accountWebSocket') accountWebSocket: $WebSocket;
  @Input('boxClientCount') boxClientCount: number;

  box: WordCloudModel;
  entries: any[];
  uniqueEntries: any[]; 
  loadingEntries: boolean;
  loadingBox: boolean;
  ws: $WebSocket;

  constructor(
    private notificationsService: NotificationsService,
    private uniqueEntriesPipe: UniqueEntriesPipe,
    private route: ActivatedRoute,
    private wordCloudService: WordCloudService,
    private webSocketService: WebSocketService,
    private processingService: ProcessingService
  ) {
    this.box = {
      BaseBox: {}
    };
    this.uniqueEntries = [];
    this.entries = [];
  }

  toggleEntry(entry: any) {
    this.toggleWordCloudEntryVisibility(entry);
  }

  hideAllEntries() {
    this.hideAllWordCloudEntries();
  }

  unhideAllEntries() {
    this.unhideAllWordCloudEntries();
  }

  private hideAllWordCloudEntries() {
    let wordCloudModel: WordCloudModel = {
      BoxId: this.box.BaseBox.ID,
    }
    this.processingService.enableProcessingAnimation();
    this.wordCloudService.hideAllEntries(wordCloudModel)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("Could not hide all entries.");
        return Observable.throw(error);
      })
      .subscribe(data => {
        this.processingService.disableProcessingAnimation();
        let wordcloudString = JSON.stringify(wordCloudModel);
        let message: WebsocketMessageModel = { MessageType: "WordCloudResponseHideAll", MessageData: wordcloudString }
        this.ws.send(message);
        this.uniqueEntries.forEach(entry => {
          entry.IsHidden = true;
        });
        this.notificationsService.openSnackBarDefault("All entries are now hidden.");
      });
  }

  private unhideAllWordCloudEntries() {
    let wordCloudModel: WordCloudModel = {
      BoxId: this.box.BaseBox.ID,
    }
    this.processingService.enableProcessingAnimation();
    this.wordCloudService.unhideAllEntries(wordCloudModel)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("Could not unhide all entries.");
        return Observable.throw(error);
      })
      .subscribe(data => {
        this.processingService.disableProcessingAnimation();
        let wordcloudResponsesString = JSON.stringify(this.entries);
        let message: WebsocketMessageModel = { MessageType: "WordCloudResponseUnhideAll", MessageData: wordcloudResponsesString }
        this.ws.send(message);
        this.uniqueEntries.forEach(entry => {
          entry.IsHidden = false;
        });
        this.notificationsService.openSnackBarDefault("All entries are now displayed.");
      });
  }

  private toggleWordCloudEntryVisibility(entry: any) {
    this.processingService.enableProcessingAnimation();
    if (entry.IsHidden) {
      this.wordCloudService.unhideEntry(this.box.BaseBox.Code, entry)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          console.error("Could not unhide word cloud entry visibility.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          entry.IsHidden = !entry.IsHidden;
          let entryString = JSON.stringify(entry);
          let message: WebsocketMessageModel = { MessageType: "WordCloudResponseUnhide", MessageData: entryString }
          this.ws.send(message);
        });
    } else {
      this.wordCloudService.hideEntry(this.box.BaseBox.Code, entry)
        .catch(error => {
          this.processingService.disableProcessingAnimation();
          console.error("Could not hide word cloud entry visibility.", error);
          return Observable.throw(error);
        })
        .subscribe(data => {
          this.processingService.disableProcessingAnimation();
          entry.IsHidden = !entry.IsHidden;
          let entryString = JSON.stringify(entry);
          let message: WebsocketMessageModel = { MessageType: "WordCloudResponseHide", MessageData: entryString }
          this.ws.send(message);
        });
    }
  }

  private initializeWordCloudDetails(code: string) {
    this.loadingEntries = true;
    this.wordCloudService.getEntries(code)
      .catch(error => {
        this.loadingEntries = false;
        return Observable.throw(error);
      })
      .subscribe(entries => {
        this.loadingEntries = false;
        this.entries = entries;
        this.uniqueEntries = this.uniqueEntriesPipe.transform(entries);
      });
    this.ws = this.boxWebSocket;
    this.ws.onMessage(this.onWordCloudHubMessage, { autoApply: false });
  }

  private onWordCloudHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let entry = JSON.parse(response.MessageData);
    if (response.MessageType == "WordCloudResponse") {
      this.entries.push(entry)
      this.uniqueEntries = this.uniqueEntriesPipe.transform(this.entries);
    } else if (response.MessageType == "DashboardClientCount") {
      let count = JSON.parse(response.MessageData).ClientCount;
      this.boxClientCount = count;
    } else if (response.MessageType == "WordCloudResponseHide") {
      let word = JSON.parse(response.MessageData).Response;
      let affectedEntry: any = this.entries.find(submittedEntry => submittedEntry.Response == word);
      affectedEntry.IsHidden = true;
    } else if (response.MessageType == "WordCloudResponseUnhide") {
      let word = JSON.parse(response.MessageData).Response;
      let affectedEntry: any = this.entries.find(submittedEntry => submittedEntry.Response == word);
      affectedEntry.IsHidden = false;
    } else if (response.MessageType == "WordCloudResponseHideAll") {
      this.uniqueEntries.forEach(entry => {
        entry.IsHidden = true;
      });
    } else if (response.MessageType == "WordCloudResponseUnhideAll") {
      this.uniqueEntries.forEach(entry => {
        entry.IsHidden = false;
      });
    }
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
      this.box.BaseBox.Code = code;
    });
    this.loadingBox = true;
    this.wordCloudService.getWordCloud(code)
      .catch(error => {
        console.error("Could not get base box to initialize box details.", error);
        return Observable.throw(error);
      })
      .subscribe(box => {
        this.loadingBox = false;
        this.box = box;
        this.initializeWordCloudDetails(code);
      });
  }
}
