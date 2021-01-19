import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickStatParticipantsComponent } from './quick-stat-participants.component';

describe('QuickStatParticipantsComponent', () => {
  let component: QuickStatParticipantsComponent;
  let fixture: ComponentFixture<QuickStatParticipantsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickStatParticipantsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickStatParticipantsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
