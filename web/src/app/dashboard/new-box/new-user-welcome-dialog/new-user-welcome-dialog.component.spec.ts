import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewUserWelcomeDialogComponent } from './new-user-welcome-dialog.component';

describe('NewUserWelcomeDialogComponent', () => {
  let component: NewUserWelcomeDialogComponent;
  let fixture: ComponentFixture<NewUserWelcomeDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewUserWelcomeDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewUserWelcomeDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
