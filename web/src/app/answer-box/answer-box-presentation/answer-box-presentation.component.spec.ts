import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AnswerBoxPresentationComponent } from './answer-box-presentation.component';

describe('AnswerBoxPresentationComponent', () => {
  let component: AnswerBoxPresentationComponent;
  let fixture: ComponentFixture<AnswerBoxPresentationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AnswerBoxPresentationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnswerBoxPresentationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
