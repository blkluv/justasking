import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AnswerBoxDetailsComponent } from './answer-box-details.component';

describe('AnswerBoxDetailsComponent', () => {
  let component: AnswerBoxDetailsComponent;
  let fixture: ComponentFixture<AnswerBoxDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AnswerBoxDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnswerBoxDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
