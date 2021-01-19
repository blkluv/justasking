import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUpgradePromptDialogComponent } from './account-upgrade-prompt-dialog.component';

describe('AccountUpgradePromptDialogComponent', () => {
  let component: AccountUpgradePromptDialogComponent;
  let fixture: ComponentFixture<AccountUpgradePromptDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUpgradePromptDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUpgradePromptDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
