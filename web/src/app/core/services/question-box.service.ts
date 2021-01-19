import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/timeout';
import 'rxjs/add/observable/throw';

import { BaseApiService } from './base-api.service';
import { QuestionBoxModel } from '../models/question-box.model';
import { QuestionBoxEntryModel } from '../models/question-box-entry.model';
import { QuestionBoxEntryVoteModel } from '../models/question-box-entry-vote.model';
import { QuestionBoxEntryVoteRepoService } from '../../core/repos/question-box-entry-vote-repo.service';

@Injectable()
export class QuestionBoxService {

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
    private questionBoxEntryVoteRepoService: QuestionBoxEntryVoteRepoService
  ) { }

  createQuestionBox(questionBox: QuestionBoxModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let data = questionBox;
    return this.http.post(`${baseUrl}questionbox`, data, options)
      .catch(this.baseApiService.handleError);
  }

  getQuestionBox(code: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    var questionbox = this.http.get(`${baseUrl}questionbox/code/${code}`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let questionBox = response.json() as any;
        return questionBox;
      });
    return questionbox;
  }

  vote(entry: QuestionBoxEntryVoteModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<QuestionBoxEntryVoteModel>(observer => {
      //check indexedDb for vote record on this entry
      this.questionBoxEntryVoteRepoService.getById(entry.EntryId)
        .catch(error => {
          console.error("Could not get question poll entry votes successfully", error);
          return Observable.throw(error);
        })
        .subscribe(storedVote => {
          //if record does not exist
          if (storedVote == null || typeof (storedVote) == 'undefined') {
            //perform vote
            this.insertVote(entry)
              .catch(error => {
                console.error("Could not insert question poll entry vote successfully", error);
                return Observable.throw(error);
              })
              .subscribe(response => {
                observer.next(entry);
              });
          } else {
            //if record exists
            //compare VoteTypes between entry and retrieved record
            if (storedVote.VoteType == entry.VoteType) {
              //if VoteTypes match, remove vote record from indexedDb
              if (storedVote.VoteType == "upvote") {
                entry.IsUndo = true;
                this.undoUpvote(entry)
                  .catch(error => {
                    console.error("Could not undo question poll entry upvote successfully", error);
                    return Observable.throw(error);
                  })
                  .subscribe(response => {
                    observer.next(entry);
                  });
              } else if (storedVote.VoteType == "downvote") {
                entry.IsUndo = true;
                this.undoDownvote(entry)
                  .catch(error => {
                    console.error("Could not undo question poll entry downvote successfully", error);
                    return Observable.throw(error);
                  })
                  .subscribe(response => {
                    observer.next(entry);
                  });
              }
            } else {
              //if VoteTypes do not match update the retrieved record VoteType to the entry VoteType
              if (storedVote.VoteType == "upvote") {
                entry.IsDownvoteFromUpvote = true;
                this.downvoteFromUpvote(entry)
                  .catch(error => {
                    console.error("Could not perform downvote from upvote for question poll entry", error);
                    return Observable.throw(error);
                  })
                  .subscribe(response => {
                    observer.next(entry);
                  });
              } else if (storedVote.VoteType == "downvote") {
                entry.IsUpvoteFromDownvote = true;
                this.upvoteFromDownvote(entry)
                  .catch(error => {
                    console.error("Could not perform upvote from downvote for question poll entry", error);
                    return Observable.throw(error);
                  })
                  .subscribe(response => {
                    observer.next(entry);
                  });
              }
            }
          }
        });
    });
    return response;
  }

  private insertVote(entry: QuestionBoxEntryVoteModel): Observable<QuestionBoxEntryVoteModel> {
    //make api call
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<any>(observer => {
      this.http.post(`${baseUrl}questionboxentryvote`, entry, options)
      .catch(this.baseApiService.handleError)
        .map(this.baseApiService.extractData)
        .subscribe((vote: QuestionBoxEntryVoteModel) => {
          this.questionBoxEntryVoteRepoService.insert(entry)
            .catch(error => {
              console.error("Could not insert vote for question poll entry in browser repo.", error);
              return Observable.throw(error);
            })
            .subscribe(response => {
              observer.next(entry);
            });
        });
    });
    return response;
  }

  private upvoteFromDownvote(entry: QuestionBoxEntryVoteModel): Observable<any> {
    //make api call
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<any>(observer => {
      this.http.post(`${baseUrl}questionboxentryvote/upvotefromdownvote`, entry, options)
        .catch(this.baseApiService.handleError)
        .map(this.baseApiService.extractData)
        .subscribe((vote: QuestionBoxEntryVoteModel) => {
          this.questionBoxEntryVoteRepoService.update(entry)
            .catch(error => {
              console.error("Could not update vote for question poll entry in browser repo.", error);
              return Observable.throw(error);
            })
            .subscribe(response => {
              observer.next(entry);
            });
        });
    });
    return response;
  }

  private downvoteFromUpvote(entry: QuestionBoxEntryVoteModel): Observable<any> {
    //make api call
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<any>(observer => {
      this.http.post(`${baseUrl}questionboxentryvote/downvotefromupvote`, entry, options)
        .catch(this.baseApiService.handleError)
        .map(this.baseApiService.extractData)
        .subscribe((vote: QuestionBoxEntryVoteModel) => {
          this.questionBoxEntryVoteRepoService.update(entry)
            .catch(error => {
              console.error("Could not downvote from upvote for question poll entry in browser repo.", error);
              return Observable.throw(error);
            })
            .subscribe(response => {
              observer.next(vote);
            });
        });
    });
    return response;
  }

  private undoUpvote(entry: QuestionBoxEntryVoteModel): Observable<any> {
    //make api call
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<any>(observer => {
      this.http.post(`${baseUrl}questionboxentryvoteundo/up`, entry, options)
        .catch(this.baseApiService.handleError)
        .map(this.baseApiService.extractData)
        .subscribe((vote: QuestionBoxEntryVoteModel) => {
          this.questionBoxEntryVoteRepoService.removeById(entry.EntryId)
            .catch(error => {
              console.error("Could not undo upvote for question poll entry in browser repo.", error);
              return Observable.throw(error);
            })
            .subscribe(response => {
              observer.next(vote);
            });
        });
    });
    return response;
  }

  private undoDownvote(entry: QuestionBoxEntryVoteModel): Observable<any> {
    //make api call
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<any>(observer => {
      this.http.post(`${baseUrl}questionboxentryvoteundo/down`, entry, options)
        .catch(this.baseApiService.handleError)
        .map(this.baseApiService.extractData)
        .subscribe((vote: QuestionBoxEntryVoteModel) => {
          this.questionBoxEntryVoteRepoService.removeById(entry.EntryId)
            .catch(error => {
              console.error("Could not undo downvote for question poll entry in browser repo.", error);
              return Observable.throw(error);
            })
            .subscribe(response => {
              observer.next(vote);
            });
        });
    });
    return response;
  }

  submitEntry(entry: QuestionBoxEntryModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    let response = new Observable<any>(observer => {
      this.http.post(`${baseUrl}questionboxentry`, entry, options)
        .catch(this.baseApiService.handleError)
        .map(this.baseApiService.extractData)
        .subscribe((submittedEntry: QuestionBoxEntryModel) => {
          entry.EntryId = submittedEntry.EntryId;
          let vote: QuestionBoxEntryVoteModel = {
            Code: entry.Code,
            EntryId: entry.EntryId,
            VoteType: 'upvote',
            VoteValue: 1
          }
          this.questionBoxEntryVoteRepoService.insert(vote)
            .catch(error => {
              console.error("Could not submit question poll entry in browser repo.", error);
              return Observable.throw(error);
            })
            .subscribe(response => {
              observer.next(entry);
            });
        })
    });
    return response;
  }

  hideEntry(code: string, response: QuestionBoxEntryModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}questionbox/hide/`, response, options)
      .catch(this.baseApiService.handleError);
  }

  unhideEntry(code: string, response: QuestionBoxEntryModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}questionbox/unhide/`, response, options)
      .catch(this.baseApiService.handleError);
  }

  hideAllEntries(questionBox: QuestionBoxModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}questionbox/hideall/`, questionBox, options)
      .catch(this.baseApiService.handleError);
  }

  unhideAllEntries(questionBox: QuestionBoxModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}questionbox/unhideall/`, questionBox, options)
      .catch(this.baseApiService.handleError);
  }

  getPartialEntries(questionBoxCode: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}questionbox/partial/code/${questionBoxCode}`, options)
      .map(this.baseApiService.extractData)
      .catch(this.baseApiService.handleError);
  }

  getEntries(questionBoxCode: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}questionboxentries/code/${questionBoxCode}`, options)
      .map(this.baseApiService.extractData)
      .catch(this.baseApiService.handleError);
  }

  getVisibleEntries(questionBoxCode: string): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}questionboxentriesvisible/code/${questionBoxCode}`, options)
      .map(this.baseApiService.extractData)
      .catch(this.baseApiService.handleError);
  }

  getAllUserVotesFromLocal(code: string): Observable<any> {
    return this.questionBoxEntryVoteRepoService.getAllByCode(code);
  }

}
