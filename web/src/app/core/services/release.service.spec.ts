import { TestBed, inject } from '@angular/core/testing';

import { ReleaseService } from './release.service';

describe('ReleaseService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ReleaseService]
    });
  });

  it('should be created', inject([ReleaseService], (service: ReleaseService) => {
    expect(service).toBeTruthy();
  }));
});
