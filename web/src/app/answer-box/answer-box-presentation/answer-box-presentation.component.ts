import { Component, Input, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { UUID } from 'angular2-uuid';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { RandomizerService } from '../../core/services/randomizer.service';
import { AnswerBoxService } from '../../core/services/answer-box.service';
import { WebSocketService } from '../../core/services/web-socket.service';

import { AnswerBoxModel } from '../../core/models/answer-box.model';
import { AnswerBoxEntryModel } from '../../core/models/answer-box-entry.model';
import { QuestionModel } from '../../core/models/question.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { BaseBoxModel } from '../../core/models/base-box.model';

@Component({
  selector: 'app-answer-box-presentation',
  templateUrl: './answer-box-presentation.component.html',
  styleUrls: ['./answer-box-presentation.component.scss']
})
export class AnswerBoxPresentationComponent implements OnInit {

  @Input('box') baseBox : BaseBoxModel;
  @Input('model') answerBoxModel : AnswerBoxModel;
  @Input('boxWebSocket') boxWebSocket : $WebSocket;
  
  ws: $WebSocket;
  box: AnswerBoxModel;
  isPreview: boolean;
  questionsIndexes: any;
  visibleQuestionsCount: number = 0;
  loading: boolean;
  
  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private randomizerService: RandomizerService,
    private answerBoxService: AnswerBoxService,
    private webSocketService: WebSocketService
  ) { 
      this.loading = false;
      this.isPreview = true;
      this.box = { BaseBox: {} }; 
  }
    
  private removeFromAnswerBoxEntries(entry:AnswerBoxEntryModel){
    let questionIndex = this.questionsIndexes[entry.QuestionId];
    if(typeof(this.box.Questions[questionIndex].Entries) != 'undefined'){
      this.box.Questions[questionIndex].Entries.forEach((submittedEntry,index) => {
        if(submittedEntry.EntryId == entry.EntryId){
          this.box.Questions[questionIndex].Entries.splice(index,1);
        }  
      });
    };
  }

  private addEntryToQuestion(entry:AnswerBoxEntryModel){
    let questionIndex = this.questionsIndexes[entry.QuestionId];
    if(typeof(this.box.Questions[questionIndex].Entries) == 'undefined'){
      this.box.Questions[questionIndex].Entries = [entry];
    }else{
      this.box.Questions[questionIndex].Entries.push(entry);
    }
  }
  
  addEntriesToAnswerBoxQuestion(question: QuestionModel){
    let questionIndex = this.questionsIndexes[question.QuestionId];
    this.box.Questions[questionIndex].Entries = question.Entries;
  }
  
  removeAllEntriesFromAnswerBoxQuestion(question){
    let questionIndex = this.questionsIndexes[question.QuestionId];
    this.box.Questions[questionIndex].Entries = [];
  }

  private onAnswerBoxHubMessage = (msg: MessageEvent)=> {
    let response:WebsocketMessageModel = JSON.parse(msg.data);
    if (response.MessageType == "AnswerBoxEntry"){
      let entries:AnswerBoxEntryModel[] = JSON.parse(response.MessageData);
      entries.forEach(entry => {
        this.addEntryToQuestion(entry);
      });
    } else if (response.MessageType == "AnswerBoxEntryUnhide"){
      let entry = JSON.parse(response.MessageData);
      this.addEntryToQuestion(entry);
    } else if (response.MessageType == "AnswerBoxEntryHide"){
      let entry = JSON.parse(response.MessageData);
      this.removeFromAnswerBoxEntries(entry);
    }else if (response.MessageType == "AnswerBoxQuestionUnhideAll"){
      let question = JSON.parse(response.MessageData);
      this.addEntriesToAnswerBoxQuestion(question);
    } else if (response.MessageType == "AnswerBoxQuestionHideAll"){
      let question = JSON.parse(response.MessageData);
      this.removeAllEntriesFromAnswerBoxQuestion(question);
    } else if (response.MessageType == "AnswerBoxQuestionDeactivate"){
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = broadcastedQuestion.IsActive;
          this.resetForm();
        }
      });
    } else if (response.MessageType == "AnswerBoxQuestionActivate"){
      let broadcastedQuestion = JSON.parse(response.MessageData);
      this.box.Questions.forEach(question => {
        if(question.QuestionId == broadcastedQuestion.QuestionId){
          question.IsActive = broadcastedQuestion.IsActive;
          this.resetForm();
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
    this.visibleQuestionsCount = 0;
    this.box.Questions.forEach(question => {
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
      this.isPreview = false;
      this.ws = this.boxWebSocket;
      this.ws.onMessage(this.onAnswerBoxHubMessage,{autoApply: false});
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
          this.answerBoxService.getVisibleAnswerBoxEntries(code)
            .catch(error=>{ 
                return Observable.throw(error);
            })
            .subscribe(entries=>{ 
              this.questionsIndexes = this.answerBoxService.getQuestionIndexesAssociativeArray(this.box.Questions);
              this.box.Questions = this.answerBoxService.mapQuestionsWithEntries(this.questionsIndexes, this.box.Questions, entries);
            });
        });
    }else{
      this.isPreview = true;
      this.box = this.answerBoxModel;
    } 
  }

}
