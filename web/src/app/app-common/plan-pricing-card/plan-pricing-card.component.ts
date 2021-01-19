import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { FeatureService } from '../../core/services/feature.service';
import { PricingPlanCardModel } from '../../core/models/pricing-plan-card.model';

import { PlanTypes } from '../../core/constants/planTypes.constants';

@Component({
  selector: 'app-plan-pricing-card',
  templateUrl: './plan-pricing-card.component.html',
  styleUrls: ['./plan-pricing-card.component.scss']
})
export class PlanPricingCardComponent implements OnInit {

  @Input("model") pricingPlanCardModel: PricingPlanCardModel;
  planTypes:any;
  
  constructor(
    private router: Router,
    private featureService: FeatureService
  ) {
    this.planTypes = PlanTypes;
   }

  confirm() {
    this.pricingPlanCardModel.ConfirmAction(this.pricingPlanCardModel.SelectetdDuration)
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe(response => {
      })
  }

  ngOnInit() {
    if (!this.pricingPlanCardModel) {
      this.pricingPlanCardModel = {
        PlanType: this.planTypes.BASIC,
        ConfirmText: "Start now",
        FeatureSections: this.featureService.getBasicFeatureSections(),
        ConfirmAction: () => new Observable<any>(observable => {
          this.router.navigate(["login"], { queryParams: { plan: this.planTypes.BASIC } })
          observable.next();
        })
      };
    } else {
      if (this.pricingPlanCardModel.PremiumDurations) {
        this.pricingPlanCardModel.SelectetdDuration =
          this.pricingPlanCardModel.PremiumDurations
          && this.pricingPlanCardModel.PremiumDurations.length > 0
          && this.pricingPlanCardModel.PremiumDurations[0];

        this.pricingPlanCardModel.PremiumDurations.forEach(duration => {
          if (duration.Selected) {
            this.pricingPlanCardModel.SelectetdDuration = duration;
            return;
          }
        });
      }
    }
  }

}
