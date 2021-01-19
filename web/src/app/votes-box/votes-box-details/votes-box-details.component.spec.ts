import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VotesBoxDetailsComponent } from './votes-box-details.component';

describe('VotesBoxDetailsComponent', () => {
  let component: VotesBoxDetailsComponent;
  let fixture: ComponentFixture<VotesBoxDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VotesBoxDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VotesBoxDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
