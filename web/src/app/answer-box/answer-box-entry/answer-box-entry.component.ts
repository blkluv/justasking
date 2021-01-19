import { Component, Input, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { UUID } from 'angular2-uuid';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { AnswerBoxModel } from '../../core/models/answer-box.model';
import { AnswerBoxEntryModel } from '../../core/models/answer-box-entry.model';
import { QuestionModel } from '../../core/models/question.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

import { AnswerBoxService } from '../../core/services/answer-box.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-answer-box-entry',
  templateUrl: './answer-box-entry.component.html',
  styleUrls: ['./answer-box-entry.component.scss']
})
export class AnswerBoxEntryComponent implements OnInit {

  @Input('box') baseBox: BaseBoxModel;
  @Input('model') answerBoxModel : AnswerBoxModel;
  @Input('boxWebSocket') boxWebSocket : $WebSocket;

  box: AnswerBoxModel;
  formSubmitted: boolean;
  ws: $WebSocket;
  visibleQuestionsCount: number = 0;
  loading: boolean;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private answerBoxService: AnswerBoxService,
    private webSocketService: WebSocketService
  ) { 
      this.loading = false;
      this.formSubmitted = false;
      this.box = { BaseBox: {} };
  }

  onSubmit(){
    let entries:AnswerBoxEntryModel[] = [];
    this.box.Questions.forEach(question => {
      if(question.IsActive){
        entries.push({QuestionId: question.QuestionId, Entry: question.Entry, IsHidden: false});
      }
    });
    this.answerBoxService.submitEntries(entries) 
      .catch(error=>{ 
        return Observable.throw(error);
      })
      .subscribe(entries=>{
        this.formSubmitted = true; 
        let entriesString = JSON.stringify(entries);
        let message:WebsocketMessageModel = { MessageType:"AnswerBoxEntry", MessageData: entriesString}
        this.ws.send(message);   
    });
  }

  allowSubmitAgain(){
    this.box.Questions.forEach((question:QuestionModel) => {
      question.Entry = null;
    });
    this.formSubmitted = false;
  }
  
  private onAnswerBoxHubMessage = (msg: MessageEvent)=> {
    let response:WebsocketMessageModel = JSON.parse(msg.data);
    if (response.MessageType == "AnswerBoxQuestionDeactivate"){
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = broadcastedQuestion.IsActive;
          this.resetForm()
        }
      });
    } else if (response.MessageType == "AnswerBoxQuestionActivate"){
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = broadcastedQuestion.IsActive;
          this.resetForm()
        }
      });
    } else if (response.MessageType == "OpenBox"){
      let box = JSON.parse(response.MessageData); 
      this.box.BaseBox.PhoneNumber = box.PhoneNumber;
    } else if (response.MessageType == "CloseBox"){
      let box = JSON.parse(response.MessageData); 
      this.box.BaseBox.PhoneNumber = box && box.PhoneNumber;
    }
  }

  private resetForm(){
    this.formSubmitted = false;
    this.visibleQuestionsCount = 0;
    this.box.Questions.forEach(question => {
      question.Entry = "";
      if(question.IsActive){
        this.visibleQuestionsCount++;
      }
    });
  }

  ngOnInit() {
    let code : string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });
    if(!this.answerBoxModel.IsPreview){ 
      this.box.BaseBox = this.baseBox;
      this.loading = true;
      this.answerBoxService.getAnswerBox(code)
        .catch(error=>{ 
          this.loading = false;
          console.error("Could not get answer poll.", error);
          return Observable.throw(error);
        })
        .subscribe(answerbox => {
          this.loading = false;
          this.box = answerbox;
          this.box.Questions.forEach(question => {
            if(question.IsActive){
              this.visibleQuestionsCount++;
            }
          });
        });
      this.ws = this.boxWebSocket;
      this.ws.onMessage(this.onAnswerBoxHubMessage,{autoApply: false});
    }else{
      this.box = this.answerBoxModel;
      this.resetForm();
    } 
  }

}
