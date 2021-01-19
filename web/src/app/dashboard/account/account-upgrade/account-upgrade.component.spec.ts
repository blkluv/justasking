import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUpgradeComponent } from './account-upgrade.component';

describe('AccountUpgradeComponent', () => {
  let component: AccountUpgradeComponent;
  let fixture: ComponentFixture<AccountUpgradeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUpgradeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUpgradeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
