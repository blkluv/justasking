import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AlreadyUpgradedComponent } from './already-upgraded.component';

describe('AlreadyUpgradedComponent', () => {
  let component: AlreadyUpgradedComponent;
  let fixture: ComponentFixture<AlreadyUpgradedComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AlreadyUpgradedComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AlreadyUpgradedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
