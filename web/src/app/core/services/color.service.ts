import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';

import { BaseApiService } from './base-api.service';
import { ColorModel } from '../models/color.model';

@Injectable()
export class ColorService {

  colors: ColorModel[];

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
  ) { }

  getColors(): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    var colors = this.http.get(`${baseUrl}themes`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let themes = response.json() as any;
        return themes;
      });
    return colors;
  }

  getTheme(themeId: number): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    var themeInfo = this.http.get(`${baseUrl}theme/${themeId}`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let theme = response.json() as any;
        return theme;
      });
    return themeInfo;
  }
}
