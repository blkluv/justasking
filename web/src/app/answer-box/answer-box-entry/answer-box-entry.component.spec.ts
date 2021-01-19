import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AnswerBoxEntryComponent } from './answer-box-entry.component';

describe('AnswerBoxEntryComponent', () => {
  let component: AnswerBoxEntryComponent;
  let fixture: ComponentFixture<AnswerBoxEntryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AnswerBoxEntryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnswerBoxEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
