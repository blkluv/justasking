import { TestBed, inject } from '@angular/core/testing';

import { BaseRepoService } from './base-repo.service';

describe('BaseRepoService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [BaseRepoService]
    });
  });

  it('should be created', inject([BaseRepoService], (service: BaseRepoService) => {
    expect(service).toBeTruthy();
  }));
});
