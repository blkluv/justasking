import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUpgradedDialogComponent } from './account-upgraded-dialog.component';

describe('AccountUpgradedDialogComponent', () => {
  let component: AccountUpgradedDialogComponent;
  let fixture: ComponentFixture<AccountUpgradedDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUpgradedDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUpgradedDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
