import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BoxDetailsComponent } from './box-details.component';

describe('BoxDetailsComponent', () => {
  let component: BoxDetailsComponent;
  let fixture: ComponentFixture<BoxDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BoxDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BoxDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
