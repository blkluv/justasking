import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { JoinAccountComponent } from './join-account.component';

describe('JoinAccountComponent', () => {
  let component: JoinAccountComponent;
  let fixture: ComponentFixture<JoinAccountComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ JoinAccountComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(JoinAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
