import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/throw';

import { BaseApiService } from '../../../core/services/base-api.service';
import { StripeService } from '../../../core/services/stripe.service';
import { UserService } from '../../../core/services/user.service';
import { PlanService } from '../../../core/services/plan.service';
import { AccountService } from '../../../core/services/account.service';
import { ProcessingService } from '../../../core/services/processing.service';

import { FeatureService } from '../../../core/services/feature.service';

import { UserModel } from '../../../core/models/user.model';
import { PlanModel } from '../../../core/models/plan.model';
import { PricingPlanCardModel } from '../../../core/models/pricing-plan-card.model';
import { PlanTypes } from '../../../core/constants/planTypes.constants';

@Component({
  selector: 'app-account-upgrade',
  templateUrl: './account-upgrade.component.html',
  styleUrls: ['./account-upgrade.component.scss']
})
export class AccountUpgradeComponent implements OnInit {

  plans: PlanModel[];
  premiumPlanCardModel: PricingPlanCardModel;

  constructor(
    private router: Router,
    private baseApiService: BaseApiService,
    private stripeService: StripeService,
    public userService: UserService,
    private planService: PlanService,
    private accountService: AccountService,
    private processingService: ProcessingService,
    private featureService: FeatureService
  ) {
    this.plans = [];
  }

  openCheckout(duration: string) {
    let planName = "";

    if (duration == '1 Week') {
      planName = 'PREMIUM-WEEK';
    } else if (duration == '1 Month') {
      planName = 'PREMIUM-MONTH';
    } else if (duration == '1 Year') {
      planName = 'PREMIUM-YEAR';
    } else {
      planName = 'PREMIUM-MONTH';
    }

    this.planService.getPlanByName(planName)
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe((plan: PlanModel) => {
        this.stripeService.openCheckout(plan, this.userService.user, (response) => {
          this.processingService.enableProcessingAnimation();
          this.stripeService.updateUsersDefaultCardOnStripe(response)
            .catch(error => {
              return Observable.throw(error);
            })
            .subscribe(response => {
              this.accountService.updateAccountPlan(plan)
                .catch(error => {
                  return Observable.throw(error);
                })
                .subscribe((newToken: string) => {
                  this.baseApiService.setToken(newToken);
                  this.baseApiService.updatedToken(newToken);
                  this.userService.getUserFromApi()
                    .catch(error => {
                      return Observable.throw(error);
                    })
                    .subscribe(() => {
                      this.router.navigate(["/dashboard/boxes"], { queryParams: { upgraded: true } });
                    })
                });
            })
        });
      });
  }

  ngOnInit() {
    this.userService.getUser()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe();

    this.planService.getPlans()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe((plans: PlanModel[]) => {
        this.plans = plans;
      });

    this.premiumPlanCardModel = {
      PlanType: PlanTypes.PREMIUM,
      ConfirmText: "Unlock these features",
      FeatureSections: this.featureService.getPremiumFeatureSections(),
      PremiumDurations: [
        { Duration: "1 Week", Price: 100 },
        { Duration: "1 Month", Price: 150, Selected: true },
        { Duration: "1 Year", Price: 750 },
      ],
      ConfirmAction: (premiumPricingPlanModel: any) => new Observable<any>(observable => {
        this.openCheckout(premiumPricingPlanModel.Duration);
        observable.next();
      })
    };
  }
}
