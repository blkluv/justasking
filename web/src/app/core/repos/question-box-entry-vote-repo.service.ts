import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';
import 'rxjs/add/observable/throw';

import { QuestionBoxEntryVoteModel } from '../models/question-box-entry-vote.model';
import { BaseRepoService } from './base-repo.service';

@Injectable()
export class QuestionBoxEntryVoteRepoService {
 
  constructor(private baseRepoService: BaseRepoService) { 
  }

  insert(vote:QuestionBoxEntryVoteModel): Observable<QuestionBoxEntryVoteModel>{
    let response = new Observable<QuestionBoxEntryVoteModel>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.QuestionBoxEntryVote], 'readwrite');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.QuestionBoxEntryVote);
        let addRequest = store.add(vote);
        this.baseRepoService.setDefaultErrorHandler(addRequest);
        addRequest.onsuccess = (e:any)=>{
          observable.next(e.target.result);
        }
      }));
    });

    return response;
  }

  getAllByCode(code:string): Observable<QuestionBoxEntryVoteModel[]>{
    let response = new Observable<QuestionBoxEntryVoteModel[]>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let votes: QuestionBoxEntryVoteModel[] = [];
        let totalCount: 0; 
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.QuestionBoxEntryVote], 'readonly');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.QuestionBoxEntryVote);
        let index = store.index('Code');
        let range = IDBKeyRange.only(code);
        let countRequest = index.count(range);
        countRequest.onsuccess = (countEvent:any)=>{
          totalCount = countEvent.target.result;
          let cursorRequest = index.openCursor(range);
          this.baseRepoService.setDefaultErrorHandler(cursorRequest);
          cursorRequest.onsuccess = (e:any)=>{
            let result = e.target.result;
            let vote = result && result.value;
            if(vote){
              votes.push(result.value);
            }
            if(votes.length === totalCount){
              observable.next(votes);
            }else{
              result.continue();
            }
          }
        }
      }));
      
    })
    return response;
  }

  getById(id:string): Observable<QuestionBoxEntryVoteModel>{
    let response = new Observable<QuestionBoxEntryVoteModel>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let vote: QuestionBoxEntryVoteModel;
        if(id && id.length > 0){
          let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.QuestionBoxEntryVote], 'readonly');
          let index = transaction.objectStore(this.baseRepoService.db.StoreNames.QuestionBoxEntryVote).index('EntryId');
          let getRequest = index.get(id);
          this.baseRepoService.setDefaultErrorHandler(getRequest);
          getRequest.onsuccess = (e:any)=>{
            vote = e.target.result;
            observable.next(vote);
          }
        }else{
          observable.next(null);
        }
      }));
    })
    return response;
  }

  update(vote: QuestionBoxEntryVoteModel): Observable<QuestionBoxEntryVoteModel>{
    let response = new Observable<QuestionBoxEntryVoteModel>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.QuestionBoxEntryVote], 'readwrite');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.QuestionBoxEntryVote);
        let index = store.index('EntryId');
        let getRequest = index.getKey(vote.EntryId);
        this.baseRepoService.setDefaultErrorHandler(getRequest);
        getRequest.onsuccess = (e:any)=>{
          let key = e.target.result;
          let putRequest = store.put(vote, key);
          this.baseRepoService.setDefaultErrorHandler(putRequest);
          putRequest.onsuccess = (e:any)=>{
            observable.next(vote);
          }
        }
      })); 
    })
    return response;
  }

  removeById(id:string): Observable<QuestionBoxEntryVoteModel>{
    let response = new Observable<QuestionBoxEntryVoteModel>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.QuestionBoxEntryVote], 'readwrite');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.QuestionBoxEntryVote);
        let index = store.index('EntryId');
        let getRequest = index.getKey(id);
        this.baseRepoService.setDefaultErrorHandler(getRequest);
        getRequest.onsuccess = (e:any)=>{
          let key = e.target.result;
          let putRequest = store.delete(key);
          this.baseRepoService.setDefaultErrorHandler(putRequest);
          putRequest.onsuccess = (e:any)=>{
            observable.next({});
          }
        }
      }));
    })
    return response;
  }
 
}
