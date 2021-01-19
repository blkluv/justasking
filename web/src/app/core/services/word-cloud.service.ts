import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';

import { BaseApiService } from './base-api.service';
import { WordCloudModel } from '../models/word-cloud.model';
import { WordCloudResponseModel } from '../models/word-cloud-response.model';

@Injectable()
export class WordCloudService {

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
  ) { }

  getWordCloud(code: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    var wordCloudResponse = this.http.get(`${baseUrl}wordcloud/code/${code}`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let wordCloud = response.json() as any;
        return wordCloud;
      });
    return wordCloudResponse;
  }

  createWordCloud(wordCloud: WordCloudModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = wordCloud;
    return this.http.post(`${baseUrl}wordcloud`, data, options)
      .catch(this.baseApiService.handleError);
  }

  submitEntry(response: WordCloudResponseModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}wordcloudresponse`, response, options)
      .catch(this.baseApiService.handleError);
  }

  hideEntry(code: string, response: WordCloudResponseModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}wordcloudresponse/hide/`, response, options)
      .catch(this.baseApiService.handleError);
  }

  unhideEntry(code: string, response: WordCloudResponseModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}wordcloudresponse/unhide/`, response, options)
      .catch(this.baseApiService.handleError);
  }

  hideAllEntries(wordCloudBoxModel: WordCloudModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}wordcloud/hideall/`, wordCloudBoxModel, options)
      .catch(this.baseApiService.handleError);
  }

  unhideAllEntries(wordCloudBoxModel: WordCloudModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}wordcloud/unhideall/`, wordCloudBoxModel, options)
      .catch(this.baseApiService.handleError);
  }

  getPartialEntries(wordCloudCode: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}wordcloudresponse/partial/code/${wordCloudCode}`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  getEntries(wordCloudCode: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}wordcloudresponse/full/code/${wordCloudCode}`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }
}
