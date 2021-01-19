import { TestBed, inject } from '@angular/core/testing';

import { VotesBoxService } from './votes-box.service';

describe('VotesBoxService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [VotesBoxService]
    });
  });

  it('should be created', inject([VotesBoxService], (service: VotesBoxService) => {
    expect(service).toBeTruthy();
  }));
});
