import { TestBed, inject } from '@angular/core/testing';

import { QuestionBoxEntryVoteRepoService } from './question-box-entry-vote-repo.service';

describe('QuestionBoxEntryVoteRepoService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [QuestionBoxEntryVoteRepoService]
    });
  });

  it('should be created', inject([QuestionBoxEntryVoteRepoService], (service: QuestionBoxEntryVoteRepoService) => {
    expect(service).toBeTruthy();
  }));
});
