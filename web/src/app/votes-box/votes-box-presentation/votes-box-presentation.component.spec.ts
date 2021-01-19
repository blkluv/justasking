import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VotesBoxPresentationComponent } from './votes-box-presentation.component';

describe('VotesBoxPresentationComponent', () => {
  let component: VotesBoxPresentationComponent;
  let fixture: ComponentFixture<VotesBoxPresentationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VotesBoxPresentationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VotesBoxPresentationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
