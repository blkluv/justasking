import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionBoxEntryComponent } from './question-box-entry.component';

describe('QuestionBoxEntryComponent', () => {
  let component: QuestionBoxEntryComponent;
  let fixture: ComponentFixture<QuestionBoxEntryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuestionBoxEntryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionBoxEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
