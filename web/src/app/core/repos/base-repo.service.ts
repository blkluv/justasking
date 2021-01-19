import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';
import 'rxjs/add/observable/throw';

import { DatabaseModel } from '../models/database.model';
 
@Injectable()
export class BaseRepoService {

  db: DatabaseModel;

  constructor() {
    this.db = {}
    this.openDb();
  }

  setDefaultErrorHandler(request) {
    if ('onerror' in request) {
      request.onerror = this.defaultErrorHandler;
    }
    if ('onblocked' in request) {
      request.onblocked = this.defaultErrorHandler;
    }
    if ('onabort' in request) {
      request.onblocked = this.defaultErrorHandler;
    }
  }

  getDbInstance(): Observable<any> {
    let response = new Observable<any>(observable => {
      if (this.db.Instance) {
        observable.next(this.db.Instance);
      } else {
        this.openDb()
          .catch(error => {
            observable.error(error);
            return Observable.throw(error);
          })
          .subscribe(database => {
            this.db.Instance = database.Instance;
            observable.next(this.db.Instance);
          });
      }
    });
    return response;
  }

  private defaultErrorHandler(e) {
    console.error(e);
  }

  //call this openDb before doing anything, if it's open, return opened db, 
  //otherwise open it.
  private openDb(): Observable<DatabaseModel> {
    let response = new Observable<DatabaseModel>(observable => {
      this.db = {
        Name: 'justaskingDB',
        Version: 2,
        Instance: {},
        StoreNames: {
          QuestionBoxEntryVote: 'QuestionBoxEntryVote',
          VotesBoxVote: 'VotesBoxVote',
        }
      }

      //opening or initiating database depending on whether it exists or not in IndexedDB
      let openRequest = window.indexedDB.open(this.db.Name, this.db.Version);
      this.setDefaultErrorHandler(openRequest);

      openRequest.onupgradeneeded = (e: any) => {
        let newVersion = e.target.result;
        if (!newVersion.objectStoreNames.contains(this.db.StoreNames.QuestionBoxEntryVote)) {
          var store = newVersion.createObjectStore(this.db.StoreNames.QuestionBoxEntryVote, { autoIncrement: true });
          store.createIndex('Code', 'Code', { unique: false });
          store.createIndex('EntryId', 'EntryId', { unique: true });
        }
        if (!newVersion.objectStoreNames.contains(this.db.StoreNames.VotesBoxVote)) {
          var store = newVersion.createObjectStore(this.db.StoreNames.VotesBoxVote, { autoIncrement: true });
          store.createIndex('Code', 'Code', { unique: false });
          store.createIndex('AnswerId', 'AnswerId', { unique: true });
          store.createIndex('QuestionId', 'QuestionId', { unique: true });
        }
      }

      openRequest.onerror = (error: any) => {
        observable.error(error);
      }

      openRequest.onsuccess = (e: any) => {
        this.db.Instance = e.target.result;
        observable.next(this.db);
      }

    });

    return response;
  }
}
