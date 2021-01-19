import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VotesBoxEntryComponent } from './votes-box-entry.component';

describe('VotesBoxEntryComponent', () => {
  let component: VotesBoxEntryComponent;
  let fixture: ComponentFixture<VotesBoxEntryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VotesBoxEntryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VotesBoxEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
