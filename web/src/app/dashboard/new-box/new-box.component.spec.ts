import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewBoxComponent } from './new-box.component';

describe('NewBoxComponent', () => {
  let component: NewBoxComponent;
  let fixture: ComponentFixture<NewBoxComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewBoxComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
