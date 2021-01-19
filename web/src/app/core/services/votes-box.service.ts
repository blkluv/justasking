import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';

import { BaseApiService } from './base-api.service';
import { VotesBoxModel } from '../models/votes-box.model';
import { VotesBoxVoteModel } from '../models/votes-box-vote.model';
import { VotesBoxQuestionModel } from '../models/votes-box-question.model';

@Injectable()
export class VotesBoxService {

  constructor(
    private http: Http,
    private baseApiService:BaseApiService
  ){}

  activateQuestion(votesBoxQuestion:VotesBoxQuestionModel):Observable<any>{
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}votesbox/activate/question/`, votesBoxQuestion, options)
    .catch(this.baseApiService.handleError);
  }
  
  deActivateQuestion(votesBoxQuestion:VotesBoxQuestionModel):Observable<any>{
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}votesbox/deactivate/question/`, votesBoxQuestion, options)
    .catch(this.baseApiService.handleError);
  }

  createVotesBox(votesBox:VotesBoxModel): Observable<any>{
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = votesBox;
    return this.http.post(`${baseUrl}votesbox`, data, options)
    .catch(this.baseApiService.handleError);
  }

  getVotesBox(code:string): Observable<any>{
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    var questionbox = this.http.get(`${baseUrl}votesbox/code/${code}`, options)
    .catch(this.baseApiService.handleError)
      .map(response => {
        let votesBox = response.json() as any;
        return votesBox;
      });
    return questionbox;
  }

  submitVotes(votes:VotesBoxVoteModel[]): Observable<any>{
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let votesSubmissionPromises: Promise<any>[] = [];
    let response: Observable<any> = new Observable<any>(observable=>{
      let submittedEntries = [];
      votes.forEach(entry => {
        let promise = this.http.post(`${baseUrl}votesbox/vote/`, entry, options)
        .catch(this.baseApiService.handleError)
        .toPromise();
        votesSubmissionPromises.push(promise);
      });

      Promise.all(votesSubmissionPromises).then((responses)=>{
        responses.forEach(response => {
          let vote = response.json();
          vote.DeletedAt = null;
          vote.CreatedBy = null;
          vote.CreatedAt = null;
          vote.UpdatedBy = null;
          vote.UpdatedAt = null;
          submittedEntries.push(vote);
        });
        observable.next(submittedEntries);
      })
    });
    
    return response;
  }  
}
