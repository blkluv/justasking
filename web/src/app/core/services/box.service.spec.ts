import { TestBed, inject } from '@angular/core/testing';

import { BoxService } from './box.service';

describe('BoxService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [BoxService]
    });
  });

  it('should ...', inject([BoxService], (service: BoxService) => {
    expect(service).toBeTruthy();
  }));
});
