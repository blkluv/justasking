import { TestBed, inject } from '@angular/core/testing';

import { AnswerBoxService } from './answer-box.service';

describe('AnswerBoxService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AnswerBoxService]
    });
  });

  it('should ...', inject([AnswerBoxService], (service: AnswerBoxService) => {
    expect(service).toBeTruthy();
  }));
});
