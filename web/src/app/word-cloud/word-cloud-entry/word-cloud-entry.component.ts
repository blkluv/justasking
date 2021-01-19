import { Component, OnInit, Input, AfterViewInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { environment } from '../../../environments/environment';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ColorModel } from '../../core/models/color.model';
import { WordCloudService } from '../../core/services/word-cloud.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { WordCloudModel } from '../../core/models/word-cloud.model';
import { WordCloudResponseModel } from '../../core/models/word-cloud-response.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { NotificationsService } from '../../core/services/notifications.service';
import { BaseBoxModel } from '../../core/models/base-box.model';

declare var WordCloud:any;

@Component({
  selector: 'app-word-cloud-entry',
  templateUrl: './word-cloud-entry.component.html',
  styleUrls: ['./word-cloud-entry.component.scss']
})
export class WordCloudEntryComponent implements OnInit {

  @Input('box') baseBox : BaseBoxModel;
  @Input('model') wordCloudModel : WordCloudModel;
  @Input('boxWebSocket') boxWebSocket : $WebSocket;
  
  wordCloudResponseModel : WordCloudResponseModel; 
  entry: string;
  box: WordCloudModel;
  ws: $WebSocket;
  loading: boolean;

  constructor(
        private notificationsService: NotificationsService,
        private route: ActivatedRoute,
        private wordCloudService: WordCloudService,
        private webSocketService: WebSocketService
  ) { 
    this.loading = false;
    this.wordCloudResponseModel = new WordCloudResponseModel();
    this.box = { BaseBox: {} };
  }

  onSubmit(){
    this.wordCloudResponseModel.Response = this.entry && this.entry.trim();
    this.wordCloudService.submitEntry(this.wordCloudResponseModel)
      .catch(error=>{  
        console.error("Could not get submit entry for word cloud.", error);
        return Observable.throw(error);
      })
      .subscribe(response => {
        let entry = response.json();
        let entryString = JSON.stringify(entry);
        let message:WebsocketMessageModel = { MessageType:"WordCloudResponse", MessageData: entryString}
        this.ws.send(message);
        this.entry = "";
        this.notificationsService.openSnackBarDefault("Thank you for your entry!");
        let entryEl = document.getElementById("boxEntry");
        entryEl.focus();
      })
  }
  
  private onWordCloudHubMessage = (msg: MessageEvent)=> {
    let response:WebsocketMessageModel = JSON.parse(msg.data);
    if (response.MessageType == "OpenBox"){
      let box = JSON.parse(response.MessageData); 
      this.box.BaseBox.PhoneNumber = box.PhoneNumber;
    } else if (response.MessageType == "CloseBox"){
      let box = JSON.parse(response.MessageData); 
      this.box.BaseBox.PhoneNumber = box && box.PhoneNumber;
    } 
  }
 
  ngOnInit() {
    let code : string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });
    if(!this.wordCloudModel.IsPreview){
      this.box.BaseBox = this.baseBox;
      this.loading = true;
      this.ws = this.boxWebSocket;
      this.ws.onMessage(this.onWordCloudHubMessage,{autoApply: false});  
      this.wordCloudService.getWordCloud(code)
        .catch(error=>{  
          this.loading = false;
          console.error("Could not get word cloud.", error);
          return Observable.throw(error);
        })
        .subscribe(wordCloud => {
          this.loading = false;
          this.wordCloudModel = wordCloud;
          this.wordCloudResponseModel.BoxId = this.wordCloudModel.BoxId;
          this.wordCloudResponseModel.Code = this.wordCloudModel.BaseBox.Code;
          this.box = this.wordCloudModel;
        });
    }else{
      this.box = this.wordCloudModel;
    }   
  }
}
