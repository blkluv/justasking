import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';
import 'rxjs/add/observable/throw';

import { VotesBoxVoteModel } from '../models/votes-box-vote.model';
import { BaseRepoService } from './base-repo.service';

@Injectable()
export class VotesBoxVoteRepoService {
  
   constructor(private baseRepoService: BaseRepoService) { 
   }
 
   insert(vote:VotesBoxVoteModel): Observable<VotesBoxVoteModel>{
     let response = new Observable<VotesBoxVoteModel>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.VotesBoxVote], 'readwrite');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.VotesBoxVote);
        let addRequest = store.add(vote);
        this.baseRepoService.setDefaultErrorHandler(addRequest);
        addRequest.onsuccess = (e:any)=>{
          observable.next(e.target.result);
        }
      })); 
     });
     return response;
   }
 
   getAllByCode(code:string): Observable<VotesBoxVoteModel[]>{
     let response = new Observable<VotesBoxVoteModel[]>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        observable.error(error);
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let votes: VotesBoxVoteModel[] = [];
        let totalCount: 0; 
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.VotesBoxVote], 'readonly');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.VotesBoxVote);
        let index = store.index('Code');
        let range = IDBKeyRange.only(code);
        let countRequest = index.count(range);
        countRequest.onerror = ()=>{observable.next([]);};
        countRequest.onsuccess = (countEvent:any)=>{
          totalCount = countEvent.target.result;
          let cursorRequest = index.openCursor(range);
          this.baseRepoService.setDefaultErrorHandler(cursorRequest);
          cursorRequest.onerror = ()=>{observable.next([]);};
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
      }))
     })
     return response;
   }
   
  getByQuestionId(id:string): Observable<VotesBoxVoteModel>{
    let response = new Observable<VotesBoxVoteModel>(observable=>{

      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let vote: VotesBoxVoteModel;
        if(id && id.length > 0){
          let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.VotesBoxVote], 'readonly');
          let index = transaction.objectStore(this.baseRepoService.db.StoreNames.VotesBoxVote).index('QuestionId');
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
  
  getByAnswerId(id:string): Observable<VotesBoxVoteModel>{
    let response = new Observable<VotesBoxVoteModel>(observable=>{ 
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let vote: VotesBoxVoteModel;
        if(id && id.length > 0){
          let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.VotesBoxVote], 'readonly');
          let index = transaction.objectStore(this.baseRepoService.db.StoreNames.VotesBoxVote).index('AnswerId');
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
 
  update(vote: VotesBoxVoteModel): Observable<VotesBoxVoteModel>{
    let response = new Observable<VotesBoxVoteModel>(observable=>{
      this.baseRepoService.getDbInstance()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((dbInstance=>{
        let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.VotesBoxVote], 'readwrite');
        let store = transaction.objectStore(this.baseRepoService.db.StoreNames.VotesBoxVote);
        let index = store.index('QuestionId');
        let getRequest = index.getKey(vote.QuestionId);
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
    });
    return response;
  }
 
   removeByQuestionId(id:string): Observable<VotesBoxVoteModel>{
     let response = new Observable<VotesBoxVoteModel>(observable=>{
        this.baseRepoService.getDbInstance()
        .catch(error=>{
          return Observable.throw(error);
        })
        .subscribe((dbInstance=>{
          let transaction = dbInstance.transaction([this.baseRepoService.db.StoreNames.VotesBoxVote], 'readwrite');
          let store = transaction.objectStore(this.baseRepoService.db.StoreNames.VotesBoxVote);
          let index = store.index('QuestionId');
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
