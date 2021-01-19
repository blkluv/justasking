import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { BaseApiService } from './base-api.service';

import { SupportIssue } from '../models/support-issue';

@Injectable()
export class SupportService {

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
  ) {
  }

  submitIssueToSupport(supportIssue: SupportIssue): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}supportissue`, supportIssue, options)
      .catch(this.baseApiService.handleError);
  }

}
