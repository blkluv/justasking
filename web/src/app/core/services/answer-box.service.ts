import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';

import { QuestionModel } from '../models/question.model';
import { AnswerBoxEntryModel } from '../models/answer-box-entry.model';
import { BaseApiService } from './base-api.service';

@Injectable()
export class AnswerBoxService {

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
  ) { }

  getAnswerBox(code: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let answerBox = this.http.get(`${baseUrl}answerbox/code/${code}`, options)
      .map(response => {
        let answerBox = response.json() as any;
        return answerBox;
      })
      .catch(this.baseApiService.handleError);
    return answerBox;
  }

  createAnswerBox(answerBox: any): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = answerBox;
    return this.http.post(`${baseUrl}answerbox`, data, options)
      .catch(this.baseApiService.handleError);
  }

  submitEntry(entry: AnswerBoxEntryModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = { Response: entry }
    return this.http.post(`${baseUrl}answerboxentry`, data, options)
      .catch(this.baseApiService.handleError);
  }

  submitEntries(entries: AnswerBoxEntryModel[]): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let entrySubmissionPromises: Promise<any>[] = [];
    let response: Observable<any> = new Observable<any>(observable => {
      let submittedEntries = [];
      entries.forEach(entry => {
        let promise = this.http.post(`${baseUrl}answerboxentry`, entry, options)
          .catch(this.baseApiService.handleError)
          .toPromise();
        entrySubmissionPromises.push(promise);
      });

      Promise.all(entrySubmissionPromises).then((responses) => {
        responses.forEach(response => {
          let entry = response.json();
          entry.DeletedAt = null;
          entry.CreatedBy = null;
          entry.CreatedAt = null;
          entry.UpdatedBy = null;
          entry.UpdatedAt = null;
          submittedEntries.push(entry);
        });
        observable.next(submittedEntries);
      })
    });

    return response;
  }

  activateQuestion(question: QuestionModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}answerbox/activate/question/`, question, options)
      .catch(this.baseApiService.handleError);
  }

  deActivateQuestion(question: QuestionModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}answerbox/deactivate/question/`, question, options)
      .catch(this.baseApiService.handleError);
  }

  hideEntry(entry: AnswerBoxEntryModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}answerbox/hide/`, entry, options)
      .catch(this.baseApiService.handleError);
  }

  unhideEntry(entry: AnswerBoxEntryModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}answerbox/unhide/`, entry, options)
      .catch(this.baseApiService.handleError);
  }

  hideAllEntries(question: QuestionModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}answerbox/hideall/question/`, question, options)
      .catch(this.baseApiService.handleError);
  }

  unhideAllEntries(question: QuestionModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}answerbox/unhideall/question/`, question, options)
      .catch(this.baseApiService.handleError);
  }

  getAnswerBoxEntries(code: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}answerboxentries/code/${code}`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let entries = response.json() as any;
        return entries;
      });
  }

  getVisibleAnswerBoxEntries(answerBoxCode: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}answerboxentriesvisible/code/${answerBoxCode}`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  getQuestionIndexesAssociativeArray(questions: QuestionModel[]): any {
    let idToIndexAssossiativeArray: any = {};

    //associate questionIds with question indexes 
    questions.forEach((question, index) => {
      idToIndexAssossiativeArray[question.QuestionId] = index;
    });

    return idToIndexAssossiativeArray;
  }

  mapQuestionsWithEntries(idToIndexAssossiativeArray: any, questions: QuestionModel[], entries: AnswerBoxEntryModel[]): QuestionModel[] {
    //assign every entry to its corresponding question
    entries.forEach(entry => {
      let questionsIndex = idToIndexAssossiativeArray[entry.QuestionId];
      if (typeof (questions[questionsIndex].Entries) == 'undefined') {
        questions[questionsIndex].Entries = [entry];
      } else {
        questions[questionsIndex].Entries.push(entry);
      }
    });

    return questions;
  }
}
