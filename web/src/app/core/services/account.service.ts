import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import { GoogleAnalyticsService } from './google-analytics.service';
import { BaseApiService } from './base-api.service';
import { AccountModel } from '../models/account.model';
import { PlanModel } from '../models/plan.model';
import { UserModel } from '../models/user.model';
import { AccountUserModel } from '../models/account-user.model';
import { AccountUserInviteModel } from '../models/account-user-invite.model';
import { CustomPlanLicenseModel } from '../models/custom-plan-license.model';

import { Roles } from '../constants/roles.constants';

@Injectable()
export class AccountService {

  accounts: AccountUserModel[];

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
    private googleAnalyticsService: GoogleAnalyticsService
  ) {
    this.accounts = [];
  }

  updateAccount(account: AccountModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}account`, account, options)
      .catch(this.baseApiService.handleError);
  }

  updateAccountPlan(plan: PlanModel): Observable<any> {
    if (plan.Name == "PREMIUM-MONTH") {
      this.googleAnalyticsService.trackPremiumUpgradeMonth();
    } else if (plan.Name == "PREMIUM-YEAR") {
      this.googleAnalyticsService.trackPremiumUpgradeYear();
    }
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}plan`, plan, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let newToken = response.json() as string;
        return newToken;
      });
  }

  updateAccountToCustomPlan(customPlanLicenseModel: CustomPlanLicenseModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}plan/custom`, customPlanLicenseModel, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let newToken = response.json() as string;
        return newToken;
      });
  }

  getAllUsersForAccount(account: AccountModel): Observable<AccountUserModel[]> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}account/users`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let users = response.json() as AccountUserModel[];
        return users;
      });
  }

  inviteUserToAccount(invite: AccountUserModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}account/invite`, invite, options)
      .catch(this.baseApiService.handleError);
  }

  resendUserAccountInvitation(invite: AccountUserModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}account/invite/resend`, invite, options)
      .catch(this.baseApiService.handleError);
  }

  getAccountInvitationByKey(inviteKey: string): Observable<AccountUserInviteModel> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}account/invite/${inviteKey}`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let accountInvitation = response.json() as AccountUserInviteModel;
        return accountInvitation;
      });
  }

  joinAccount(invitationCode: string): Observable<any> {
    let invitation = { invitationCode: invitationCode }
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}account/join`, invitation, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let newToken = response.json() as string;
        return newToken;
      });
  }

  getAccountsForUsers(): Observable<AccountUserModel[]> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}accounts`, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let accounts = response.json() as AccountUserModel[];
        this.accounts = accounts;
        return this.accounts;
      });
  }

  changeCurrentAccount(accountId: string): Observable<any> {
    let account = { accountId: accountId }
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}account/current`, account, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let newToken = response.json() as string;
        return newToken;
      });
  }

  cancelAccountInvitation(accountInvitation: AccountUserInviteModel) {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}account/invite/cancel`, accountInvitation, options)
      .catch(this.baseApiService.handleError);
  }

  cancelAccountPlan(): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}plan/cancel`, null, options)
      .catch(this.baseApiService.handleError);
  }

  updateAccountUser(accountUser: AccountUserModel): Observable<AccountUserModel> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}account/user`, accountUser, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let accountUser = response.json() as AccountUserModel;
        return accountUser;
      });
  }

  removeAccountUser(accountUser: AccountUserModel): Observable<AccountUserModel> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.put(`${baseUrl}account/user/delete`, accountUser, options)
      .catch(this.baseApiService.handleError);
  }

  transferAccountOwnership(accountUser: AccountUserModel): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}account/transfer`, accountUser, options)
      .catch(this.baseApiService.handleError)
      .map(response => {
        let newToken = response.json() as string;
        return newToken;
      });
  }

  getAssignableRoles(): any[] {
    let assignableRoles = [Roles.PRESENTER, Roles.ADMIN];
    return assignableRoles;
  }

}