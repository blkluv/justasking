import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { MatSelectChange, MatSlideToggleChange } from '@angular/material';
import { Location } from '@angular/common';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/observable/throw';

import { BaseApiService } from '../../core/services/base-api.service';
import { UserService } from '../../core/services/user.service';
import { AccountService } from '../../core/services/account.service';
import { PlanService } from '../../core/services/plan.service';
import { StripeService } from '../../core/services/stripe.service';
import { ProcessingService } from '../../core/services/processing.service';

import { PlanModel } from '../../core/models/plan.model';
import { UserModel } from '../../core/models/user.model';

@Component({
  selector: 'app-new-user',
  templateUrl: './new-user.component.html',
  styleUrls: ['./new-user.component.scss'],
})
export class NewUserComponent implements OnInit {

  plans: PlanModel[];
  plan: PlanModel;

  constructor(
    private baseApiService: BaseApiService,
    public userService: UserService,
    private accountService: AccountService,
    private planService: PlanService,
    private stripeService: StripeService,
    private router: Router,
    private route: ActivatedRoute,
    private location: Location,
    private processingService: ProcessingService,
  ) {
    this.plans = [];
    this.plan = {};
  }

  planChange(changeEvent: MatSelectChange) {
    let newlySelectedPlan: PlanModel = changeEvent.value;
    this.location.go(`/login/new?plan=${newlySelectedPlan.DisplayName}`);
  }

  continueWithPlan() {
    let account = this.userService.user.Account;
    this.accountService.updateAccount(account)
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe(response => {
        if (this.plan.Price > 0) {
          this.stripeService.openCheckout(this.plan, this.userService.user, (response) => {
            this.processingService.enableProcessingAnimation();
            this.stripeService.updateUsersDefaultCardOnStripe(response)
              .catch(error => {
                return Observable.throw(error);
              })
              .subscribe(response => {
                this.accountService.updateAccountPlan(this.plan)
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
                        this.router.navigate(["/dashboard/new"], { queryParams: { welcome: true } });
                      })
                  });
              })
          });
        } else {
          this.router.navigate(["/dashboard/new"], { queryParams: { welcome: true } });
        }
      });
  }

  ngOnInit() {
    let preselectedPlanName = this.route.snapshot.queryParams['plan'] || 'BASIC';
    let duration = this.route.snapshot.queryParams['duration'];
    this.userService.getUser()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe(() => {
        this.planService.getPlans()
          .catch(error => {
            return Observable.throw(error);
          })
          .subscribe((plans: PlanModel[]) => {
            this.plans = plans;
            let selectedPlanName = preselectedPlanName;

            if (preselectedPlanName == 'PREMIUM') {
              if (duration == '1 Week') {
                selectedPlanName = 'PREMIUM-WEEK';
              } else if (duration == '1 Month') {
                selectedPlanName = 'PREMIUM-MONTH';
              } else if (duration == '1 Year') {
                selectedPlanName = 'PREMIUM-YEAR';
              } else {
                selectedPlanName = 'PREMIUM-MONTH';
              }
            }
            this.planService.getPlanByName(selectedPlanName)
              .catch(error => {
                return Observable.throw(error);
              })
              .subscribe((plan: PlanModel) => {
                this.plan = plan;
              });
          });
      });
  }

}
