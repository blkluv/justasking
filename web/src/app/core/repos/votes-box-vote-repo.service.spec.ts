import { TestBed, inject } from '@angular/core/testing';

import { VotesBoxVoteRepoService } from './votes-box-vote-repo.service';
import { VotesBoxVoteModel } from '../models/votes-box-vote.model';

describe('VotesBoxVoteRepoService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [VotesBoxVoteRepoService]
    });
  });
  
  it('should be created', inject([VotesBoxVoteRepoService], (service: VotesBoxVoteRepoService) => {
    expect(service).toBeTruthy();
  })); 

});
