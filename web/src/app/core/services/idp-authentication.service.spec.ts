import { TestBed, inject } from '@angular/core/testing';

import { IdpAuthenticationService } from './idp-authentication.service';

describe('IdpAuthenticationService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [IdpAuthenticationService]
    });
  });

  it('should ...', inject([IdpAuthenticationService], (service: IdpAuthenticationService) => {
    expect(service).toBeTruthy();
  }));
});
