import { Injectable } from '@angular/core';

import { BaseApiService } from './base-api.service';
import { Http, Response } from '@angular/http';
import { Headers, RequestOptions } from '@angular/http';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';
import 'rxjs/add/observable/throw';

import { UserModel } from '../models/user.model';
import { OperationResultModel } from '../models/operation-result.model';

@Injectable()
export class UserService {

  public user: UserModel;
  private userPromiseContainer: Promise<UserModel>[];

  constructor(
    private http: Http,
    private router: Router,
    private baseApiService: BaseApiService
  ) {
    this.userPromiseContainer = [];
  }

  userIsLoggedIn(): boolean {
    return this.baseApiService.getToken() && true;
  }

  saveUser(user: UserModel) {
    this.user = user;
  }

  resetUser() {
    this.user = null;
    this.userPromiseContainer = [];
  }

  getUserFromApi(): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let userPromise = this.http.get(`${baseUrl}user`, options)
      .catch(this.baseApiService.handleError)
      .map((response: any) => {
        if (response.ok && response._body) {
          let data = response.json() as UserModel;
          this.saveUser(data);
          return data;
        } else {
          this.invalidateUser();
          return null;
        }
      })
      .catch(error => {
        this.invalidateUser();
        return Observable.throw(error);
      })
    return userPromise;
  }

  getUser(): Observable<UserModel> {
    let userObservable = new Observable<UserModel>();
    if (this.user != null && typeof (this.user) != 'undefined') {
      userObservable = new Observable<UserModel>(observer => {
        observer.next(this.user);
      });
    } else {
      if (this.userPromiseContainer.length > 0) {
        userObservable = new Observable<UserModel>(observer => {
          Promise.all(this.userPromiseContainer)
            .then((usersResponse: UserModel[]) => {
              this.userPromiseContainer = [];
              if (usersResponse.length > 0) {
                let user = usersResponse[0];
                this.saveUser(user);
              }
              observer.next(this.user);
            });
        });
      } else {
        let userPromise = this.getUserFromApi().toPromise() as Promise<UserModel>;
        userObservable = new Observable<UserModel>(observer => {
          userPromise.then((user: UserModel) => {
            if (user) {
              this.saveUser(user);
              observer.next(this.user);
            }
          });
          this.userPromiseContainer.push(userPromise);
        });
      }
    }
    return userObservable;
  }

  invalidateUser() {
    this.resetUser();
    this.baseApiService.removeToken();
    this.router.navigate(['/login']);
  }

  logout(): Observable<any> {
    let data = "<(｡◕‿‿◕｡)>";
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}user/logout`, data, options)
      .catch(this.baseApiService.handleError);
  }
}