import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';
import 'rxjs/add/observable/throw';

import { StripeService } from '../../../core/services/stripe.service';
import { NotificationsService } from '../../../core/services/notifications.service';
import { ProcessingService } from '../../../core/services/processing.service';
import { AccountService } from '../../../core/services/account.service';

import { UserService } from '../../../core/services/user.service';
import { FeatureService } from '../../../core/services/feature.service';

import { FeatureSectionModel } from '../../../core/models/feature-section.model';
import { UserStripeModel } from '../../../core/models/user-stripe.model';
import { UserModel } from '../../../core/models/user.model';
import { AccountModel } from '../../../core/models/account.model';

import { PlanTypes } from '../../../core/constants/planTypes.constants';

@Component({
  selector: 'app-account-details',
  templateUrl: './account-details.component.html',
  styleUrls: ['./account-details.component.scss']
})
export class AccountDetailsComponent implements OnInit {

  @Input('user') user: UserModel;

  featureSections: FeatureSectionModel[];
  userStripeModel: UserStripeModel;
  canEditAccountName: boolean;
  accountName: FormControl;
  renameAccountForm: FormGroup;
  planTypes:any;

  constructor(
    private processingService: ProcessingService,
    private notificationsService: NotificationsService,
    private stripeService: StripeService,
    private featureService: FeatureService,
    private accountService: AccountService, 
  ) {
    this.planTypes = PlanTypes;
    this.userStripeModel = {};
  }

  enableNameEditMode() {
    this.accountName.enable();
  }

  cancelAccountName() {
    this.accountName.setValue(this.user.Account.Name);
    this.accountName.disable();
  }

  saveAccountName(formValue: { Name: string }) {
    this.processingService.enableProcessingAnimation();
    let account: AccountModel = this.user.Account;
    account.Name = formValue.Name;
    this.accountService.updateAccount(account)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        return Observable.throw(error)
      })
      .subscribe(response => {
        if (response.status == 200) {
          this.accountService.getAccountsForUsers()
            .catch(error => {
              return Observable.throw(error);
            })
            .subscribe();
          this.notificationsService.openSnackBarDefault(`Account name updated!`);
        } else {
          this.notificationsService.openSnackBarError(`Account name could not be updated.`);
        }
        this.accountName.disable();
        this.processingService.disableProcessingAnimation();
      });
  }

  canRenameAccount() {
    let canManageAccount = false;
    if (this.user && this.user.RolePermissions.EditAccountName) {
      canManageAccount = true;
    }
    return canManageAccount;
  }

  ngOnInit() {
    this.featureSections = this.featureService.getFeatureSectionsFromMembershipDetails(this.user.MembershipDetails);
    this.accountName = new FormControl({ value: this.user.Account.Name, disabled: true }, [
      Validators.required,
      Validators.maxLength(50),
    ]);
    this.renameAccountForm = new FormGroup({ 'Name': this.accountName });
    this.canEditAccountName = this.canRenameAccount();
    this.stripeService.getStripeData()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe((userStripeModel: UserStripeModel) => {
        this.userStripeModel = userStripeModel;
      });
  }
}
