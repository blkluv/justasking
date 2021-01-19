import { TestBed, inject } from '@angular/core/testing';

import { QuestionBoxService } from './question-box.service';

describe('QuestionBoxService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [QuestionBoxService]
    });
  });

  it('should ...', inject([QuestionBoxService], (service: QuestionBoxService) => {
    expect(service).toBeTruthy();
  }));
});
