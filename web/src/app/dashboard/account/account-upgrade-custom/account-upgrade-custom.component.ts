import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

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
import { CustomPlanLicenseModel } from '../../../core/models/custom-plan-license.model';

import { PlanTypes } from '../../../core/constants/planTypes.constants';

@Component({
  selector: 'app-account-upgrade-custom',
  templateUrl: './account-upgrade-custom.component.html',
  styleUrls: ['./account-upgrade-custom.component.scss']
})
export class AccountUpgradeCustomComponent implements OnInit {

  plan: PlanModel;
  licenseCode: string;
  invalidCustomPlanLicense: boolean;
  customPlanCardModel: PricingPlanCardModel;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private baseApiService: BaseApiService,
    private stripeService: StripeService,
    public userService: UserService,
    private planService: PlanService,
    private accountService: AccountService,
    private processingService: ProcessingService,
    private featureService: FeatureService
  ) {
  }

  openCheckout() {
    this.stripeService.openCheckout(this.plan, this.userService.user, (response) => {
      this.processingService.enableProcessingAnimation();
      this.stripeService.updateUsersDefaultCardOnStripe(response)
        .catch(error => {
          return Observable.throw(error);
        })
        .subscribe(response => {
          let customPlanLicenseModel: CustomPlanLicenseModel = {
            AccountId: this.userService.user.Account.Id,
            PlanId: this.plan.Id,
            LicenseCode: this.licenseCode
          }
          this.accountService.updateAccountToCustomPlan(customPlanLicenseModel)
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
  }

  ngOnInit() {
    this.processingService.enableProcessingAnimation();
    this.route.params.forEach((params: Params) => {
      this.licenseCode = params["licenseCode"];
    });
    this.userService.getUser()
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        return Observable.throw(error);
      })
      .subscribe(user => {
        this.planService.getCustomPlan(this.licenseCode)
          .catch(error => {
            this.invalidCustomPlanLicense = true;
            this.processingService.disableProcessingAnimation();
            return Observable.throw(error);
          })
          .subscribe((plan: PlanModel) => {
            this.processingService.disableProcessingAnimation();
            this.plan = plan;
            this.customPlanCardModel = {
              PlanType: PlanTypes.CUSTOM,
              CustomPlan: plan,
              ConfirmText: "Unlock your custom plan",
              ConfirmAction: () => new Observable<any>(observable => {
                this.openCheckout();
                observable.next();
              })
            };
            this.customPlanCardModel.FeatureSections = this.featureService.getFeatureSectionsFromMembershipDetails(plan);
          });
      });
  }
}
