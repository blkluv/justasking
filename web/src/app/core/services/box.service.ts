import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';

import { BaseApiService } from './base-api.service';
import { UserModel } from '../models/user.model';
import { BaseBoxModel } from '../models/base-box.model';

@Injectable()
export class BoxService {


  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
  ) { }

  getRandomCode() {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}basebox/code`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  getAllBoxes(user: UserModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}basebox/account`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  getBaseBox(code: string): Observable<BaseBoxModel> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}baseboxbycode/${code}`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  getBaseBoxAuth(code: string): Observable<BaseBoxModel> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}basebox/details/${code}`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  codeIsTaken(code: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}basebox/exists/${code}`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  openBox(baseBoxId: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = "(づ｡◕‿‿◕｡)づ";
    return this.http.post(`${baseUrl}basebox/activate/${baseBoxId}`, data, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  closeBox(baseBoxId: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = "ᕙ(⇀‸↼‶)ᕗ";
    return this.http.post(`${baseUrl}basebox/deactivate/${baseBoxId}`, data, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  deleteBox(baseBoxId: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = {id:baseBoxId};
    return this.http.put(`${baseUrl}basebox/delete`, data, options)
      .catch(this.baseApiService.handleError);
  }

}
