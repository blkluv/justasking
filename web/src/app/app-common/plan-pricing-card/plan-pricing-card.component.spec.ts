import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlanPricingCardComponent } from './plan-pricing-card.component';

describe('PlanPricingCardComponent', () => {
  let component: PlanPricingCardComponent;
  let fixture: ComponentFixture<PlanPricingCardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlanPricingCardComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlanPricingCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
