import { TestBed, inject } from '@angular/core/testing';

import { FeatureRequestService } from './feature-request.service';

describe('FeatureRequestService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [FeatureRequestService]
    });
  });

  it('should be created', inject([FeatureRequestService], (service: FeatureRequestService) => {
    expect(service).toBeTruthy();
  }));
});
