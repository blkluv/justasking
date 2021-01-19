import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SupportWidgetComponent } from './support-widget.component';

describe('SupportWidgetComponent', () => {
  let component: SupportWidgetComponent;
  let fixture: ComponentFixture<SupportWidgetComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SupportWidgetComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SupportWidgetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
