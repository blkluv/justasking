import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUserListItemComponent } from './account-user-list-item.component';

describe('AccountUserListItemComponent', () => {
  let component: AccountUserListItemComponent;
  let fixture: ComponentFixture<AccountUserListItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUserListItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUserListItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
