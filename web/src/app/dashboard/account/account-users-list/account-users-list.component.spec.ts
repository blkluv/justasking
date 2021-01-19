import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUsersListComponent } from './account-users-list.component';

describe('AccountUsersListComponent', () => {
  let component: AccountUsersListComponent;
  let fixture: ComponentFixture<AccountUsersListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUsersListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUsersListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
