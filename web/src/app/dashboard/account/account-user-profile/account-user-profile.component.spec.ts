import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUserProfileComponent } from './account-user-profile.component';

describe('AccountUserProfileComponent', () => {
  let component: AccountUserProfileComponent;
  let fixture: ComponentFixture<AccountUserProfileComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUserProfileComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUserProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
