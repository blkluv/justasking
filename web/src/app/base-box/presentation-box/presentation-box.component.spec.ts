import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PresentationBoxComponent } from './presentation-box.component';

describe('PresentationBoxComponent', () => {
  let component: PresentationBoxComponent;
  let fixture: ComponentFixture<PresentationBoxComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PresentationBoxComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PresentationBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
