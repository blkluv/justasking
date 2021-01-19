import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionBoxDetailsComponent } from './question-box-details.component';

describe('QuestionBoxDetailsComponent', () => {
  let component: QuestionBoxDetailsComponent;
  let fixture: ComponentFixture<QuestionBoxDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuestionBoxDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionBoxDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
