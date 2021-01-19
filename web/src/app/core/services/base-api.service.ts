import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Headers, RequestOptions, ResponseContentType } from '@angular/http';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/observable/throw';

declare var localStorage;

import { UserService } from './user.service';

import { environment } from '../../../environments/environment';

@Injectable()
export class BaseApiService {
  apiBaseUrl: string;
  authToken: string;
  headers: Headers;
  options: RequestOptions;
  fileUploadOptions: RequestOptions;
  fileDownloadOptions: RequestOptions;

  constructor(private _http: Http, private router: Router) {
    this.apiBaseUrl = `${environment.protocol}://${environment.apiBaseDomain}/`;
    this.authToken = this.getToken();
    this.setRequestOptionsWithToken(this.authToken);
  }

  getBaseUrl(): string {
    return this.apiBaseUrl;
  }

  getRequestOptions() {
    return this.options;
  }

  getRequestFileUploadOptions() {
    return this.fileUploadOptions;
  }

  getRequestFileDownloadOptions() {
    return this.fileDownloadOptions;
  }

  clearRequestConfigs() {
    this.setRequestOptionsWithToken(null);
    this.authToken = null;
    this.options = null;
    this.fileUploadOptions = null;
    this.fileDownloadOptions = null;
  }

  getToken() {
    return localStorage.getItem('zxcv');
  }

  setToken(token: string) {
    localStorage.setItem('zxcv', token)
  }

  removeToken() {
    localStorage.removeItem('zxcv');
  }

  updatedToken(newToken: string) {
    this.authToken = newToken;
    this.setRequestOptionsWithToken(this.authToken);
  }

  setRequestOptionsWithToken(token: string) {
    let headers = new Headers(
      {
        "Content-Type": "application/json; charset=utf-8",
        "If-Modified-Since": "Mon, 26 Jul 1997 05:00:00 GMT",
        "Accept": "application/*",
        "Authorization": `Bearer ${token}`,
        "Strict-Transport-Security": environment.strictTransportSecurityValue,
        "X-Frame-Options": "deny",
        "X-XSS-Protection": "1; mode=block",
        "Cache-Control": "no-cache, no-store, max-age=0, must-revalidate",
        "Pragma": "no-cache",
        "Expires": "0",
        "X-Content-Type-Options": "nosniff"
      });
    let fileUploadHeaders = new Headers(
      {
        "Content-Type": "multipart/form-data; boundary=----WebKitFormBoundarytuwZ4AiLboBqePrQ",
        "Accept": "application/json",
        "Authorization": `Bearer ${token}`,
        "Strict-Transport-Security": environment.strictTransportSecurityValue,
        "X-Frame-Options": "deny",
        "X-XSS-Protection": "1; mode=block",
        "Cache-Control": "no-cache, no-store, max-age=0, must-revalidate",
        "Pragma": "no-cache",
        "Expires": "0",
        "X-Content-Type-Options": "nosniff"
      });
    let fileDownloadHeaders = new Headers(
      {
        "Content-Type": "application/json; charset=utf-8",
        "If-Modified-Since": "Mon, 26 Jul 1997 05:00:00 GMT",
        "Accept": "application/octet-stream",
        "Authorization": `Bearer ${token}`,
        "Strict-Transport-Security": environment.strictTransportSecurityValue,
        "X-Frame-Options": "deny",
        "X-XSS-Protection": "1; mode=block",
        "Cache-Control": "no-cache, no-store, max-age=0, must-revalidate",
        "Pragma": "no-cache",
        "Expires": "0",
        "X-Content-Type-Options": "nosniff"
      });
    this.options = new RequestOptions({ headers: headers });
    this.fileUploadOptions = new RequestOptions({ headers: fileUploadHeaders });
    this.fileDownloadOptions = new RequestOptions({ headers: fileUploadHeaders, responseType: ResponseContentType.Blob });
  }

  extractData(res: Response) {
    let obj = res.json();
    return obj || {};
  }

  handleError = (error: Response | any) => {
    if (error.status == 401) {
      this.removeToken();
      this.clearRequestConfigs()
      localStorage.removeItem('zxcv');
      this.router.navigate(['/login']);
    }
    console.error(error);
    return Observable.throw(error);
  }
}