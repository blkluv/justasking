import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickStatBoxComponent } from './quick-stat-box.component';

describe('QuickStatBoxComponent', () => {
  let component: QuickStatBoxComponent;
  let fixture: ComponentFixture<QuickStatBoxComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickStatBoxComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickStatBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
