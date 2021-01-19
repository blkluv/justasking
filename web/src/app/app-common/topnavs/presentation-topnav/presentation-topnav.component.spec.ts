import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PresentationTopnavComponent } from './presentation-topnav.component';

describe('PresentationTopnavComponent', () => {
  let component: PresentationTopnavComponent;
  let fixture: ComponentFixture<PresentationTopnavComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PresentationTopnavComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PresentationTopnavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
