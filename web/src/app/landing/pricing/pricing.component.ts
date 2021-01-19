import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { FeatureService } from '../../core/services/feature.service';
import { PricingPlanCardModel } from '../../core/models/pricing-plan-card.model';

import { PlanTypes } from '../../core/constants/planTypes.constants';

@Component({
  selector: 'app-pricing',
  templateUrl: './pricing.component.html',
  styleUrls: ['./pricing.component.scss']
})
export class PricingComponent implements OnInit {

  premiumPlanCardModel: PricingPlanCardModel;

  constructor(
    private router: Router,
    private featureService: FeatureService
  ) {
  }

  ngOnInit() {
    this.premiumPlanCardModel = {
      PlanType: PlanTypes.PREMIUM,
      ConfirmText: "Start now",
      FeatureSections: this.featureService.getPremiumFeatureSections(),
      PremiumDurations: [
        { Duration: "1 Week", Price: 100 },
        { Duration: "1 Month", Price: 150, Selected: true },
        { Duration: "1 Year", Price: 750 },
      ],
      ConfirmAction: (duration:{ Duration: string; Price: number; }) => new Observable<any>(observable => {
        this.router.navigate(["login"], { queryParams: { plan: 'PREMIUM', duration: duration.Duration } })
        observable.next();
      })
    };
  }
}
