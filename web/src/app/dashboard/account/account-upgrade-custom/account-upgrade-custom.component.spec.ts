import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountUpgradeCustomComponent } from './account-upgrade-custom.component';

describe('AccountUpgradeCustomComponent', () => {
  let component: AccountUpgradeCustomComponent;
  let fixture: ComponentFixture<AccountUpgradeCustomComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccountUpgradeCustomComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountUpgradeCustomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
