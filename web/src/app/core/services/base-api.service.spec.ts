import { TestBed, inject } from '@angular/core/testing';

import { BaseApiService } from './base-api.service';

describe('BaseApiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [BaseApiService]
    }); 
  });

  it('should ...', inject([BaseApiService], (service: BaseApiService) => {
    expect(service).toBeTruthy();
  }));
});
