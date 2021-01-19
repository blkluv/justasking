import { Component, OnInit, Input, AfterViewInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { environment } from '../../../environments/environment';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { WordCloudModel } from '../../core/models/word-cloud.model';
import { WordCloudService } from '../../core/services/word-cloud.service';
import { RandomizerService } from '../../core/services/randomizer.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-word-cloud-presentation',
  templateUrl: './word-cloud-presentation.component.html',
  styleUrls: ['./word-cloud-presentation.component.scss']
})
export class WordCloudPresentationComponent implements OnInit {

  @Input('box') baseBox : BaseBoxModel;
  @Input('model') wordCloudModel : WordCloudModel;
  @Input('boxWebSocket') boxWebSocket : $WebSocket;
  
  box: WordCloudModel;
  list: [[string,number]];
  wordCloudOptions: any;
  loading: boolean;

  constructor(
        private route: ActivatedRoute,
        private randomizerService: RandomizerService,
        private wordCloudService: WordCloudService,
        private webSocketService: WebSocketService
  ) {
    this.loading = false;
    this.box = {BaseBox: {}};
    this.list =  [["",0]];
    this.wordCloudOptions = {};
  }

  private addToWordList(word:string){
    word = word && word.trim().toLowerCase();
    let updated = false;

    this.list.forEach(function(value){
      let wordEntry = value[0] && value[0].toLowerCase();
      let wordCount = value[1];
      if(wordEntry == word){
        if(wordCount <= 60){
          wordCount = wordCount + 5;
        }else if(wordCount < 100){
          wordCount = wordCount + 1;
        }else if(wordCount <= Infinity){
          wordCount = wordCount;
        }else{
          wordCount = wordCount + 10;
        }
        value[1] = wordCount;
        updated = true;
        return;
      }
    })
    if(!updated){
      this.list.push([word,20]);
    }
  }

  private removeFromWordList(word:string){
    this.list.forEach((submittedEntry,index)=>{
      let entry = submittedEntry[0] && submittedEntry[0].trim().toLowerCase();
      word = word && word.trim().toLowerCase();
      if(entry == word){
        this.list.splice(index,1);
      }
    });
    this.list = this.list.slice(0) as [[string,number]];
  }

  private clearWordCloud(){
    this.list = [[this.box.DefaultWord,110]];
  }

  private onWordCloudHubMessage = (msg: MessageEvent)=> {
    let response:WebsocketMessageModel = JSON.parse(msg.data);
    let entry = JSON.parse(response.MessageData);
    if (response.MessageType == "WordCloudResponse") {
      this.addToWordList(entry.Response);
      this.list = this.list.slice(0) as [[string,number]];
    } else if (response.MessageType == "WordCloudResponseHide") {
      let word = JSON.parse(response.MessageData).Response;
      this.removeFromWordList(word)
    } else if (response.MessageType == "WordCloudResponseUnhide") {
      let entry = JSON.parse(response.MessageData);
      for (var index = 0; index < entry.Count; index++) {
        this.addToWordList(entry.Response);        
      }
      this.list = this.list.slice(0) as [[string,number]];
    } else if (response.MessageType == "WordCloudResponseHideAll") {
      this.clearWordCloud();
    } else if (response.MessageType == "WordCloudResponseUnhideAll") {
      let wordCloudResponses = JSON.parse(response.MessageData);
      wordCloudResponses.forEach(wordCloudResponse => {
        this.addToWordList(wordCloudResponse.Response);
      });
      this.list = this.list.slice(0) as [[string,number]]; 
    } else if (response.MessageType == "OpenBox"){
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
    this.wordCloudOptions = { 
      list: null,
      shape: "circle",
      shuffle: false,
      backgroundColor: "transparent",
      background: "transparent",
      color: "white",
      rotateRatio: 0,
      minRotation: 1,
      maxRotation: 1, 
    } 
    if(!this.wordCloudModel.IsPreview){ 
      this.box.BaseBox = this.wordCloudModel.BaseBox;
      this.loading = true;
      let ws = this.boxWebSocket;
      ws.onMessage(this.onWordCloudHubMessage,{autoApply: false}); 
      this.wordCloudService.getWordCloud(code)
        .catch(error=>{  
          this.loading = false;
          console.error("Could not get word cloud.", error);
          return Observable.throw(error);
        })
        .subscribe(wordCloud => {
          this.loading = false;
          this.wordCloudModel = wordCloud;
          this.box = wordCloud;
          this.list = [[this.box.DefaultWord,110]];
          this.wordCloudService.getEntries(code)
            .catch(error=>{  
              console.error("Could not get word cloud entries.", error);
              return Observable.throw(error);
            })
            .subscribe(cloudResponses=>{
              if(cloudResponses && cloudResponses.length > 0){
                cloudResponses.forEach(response => {
                  if(!response.IsHidden){
                    this.addToWordList(response.Response);
                  }
                });
                this.list = this.list.slice(0) as [[string,number]];
              }
            });
      });
    }else{
      let randomEntries = this.randomizerService.getRandomWordCloudEntriesWithDefault(this.wordCloudModel.DefaultWord, 75);
      this.list = randomEntries;
      this.box = this.wordCloudModel;
    }
  }
}
