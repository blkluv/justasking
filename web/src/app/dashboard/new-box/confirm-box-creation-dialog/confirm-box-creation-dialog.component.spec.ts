import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfirmBoxCreationDialogComponent } from './confirm-box-creation-dialog.component';

describe('ConfirmBoxCreationDialogComponent', () => {
  let component: ConfirmBoxCreationDialogComponent;
  let fixture: ComponentFixture<ConfirmBoxCreationDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ConfirmBoxCreationDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ConfirmBoxCreationDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
