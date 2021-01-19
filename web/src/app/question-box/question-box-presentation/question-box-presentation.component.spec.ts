import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionBoxPresentationComponent } from './question-box-presentation.component';

describe('QuestionBoxPresentationComponent', () => {
  let component: QuestionBoxPresentationComponent;
  let fixture: ComponentFixture<QuestionBoxPresentationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuestionBoxPresentationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuestionBoxPresentationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
