import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { BaseApiService } from './base-api.service';

import { FeatureRequest } from '../models/feature-request.model';

@Injectable()
export class FeatureRequestService {

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
  ) {
  }

  submitFeatureRequest(featureRequest: FeatureRequest): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}featurerequest`, featureRequest, options)
      .catch(this.baseApiService.handleError);
  }

}

