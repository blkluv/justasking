import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/observable/throw';

import { BaseApiService } from './base-api.service';
import { PlanModel } from '../models/plan.model';

@Injectable()
export class PlanService {

  plans: PlanModel[];

  constructor(
    private http: Http,
    private baseApiService: BaseApiService
  ) {
    this.plans = [];
  }

  getPlans(): Observable<PlanModel[]> {
    let response: Observable<PlanModel[]>;
    if (this.plans.length > 0) {
      response = new Observable<PlanModel[]>(observable => {
        observable.next(this.plans);
      });
    } else {
      let baseUrl = this.baseApiService.getBaseUrl();
      let options = this.baseApiService.getRequestOptions();
      response = this.http.get(`${baseUrl}priceplan`, options)
        .catch(this.baseApiService.handleError)
        .map((response: any) => {
          if (response.ok && response._body) {
            let data = response.json() as PlanModel[];
            this.plans = data;
            return data;
          } else {
            return null;
          }
        });
    }
    return response;
  }

  getCustomPlan(customPlanLicenseCode: string): Observable<PlanModel> {
    let response: Observable<PlanModel>;
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    response = this.http.get(`${baseUrl}priceplan/custom/${customPlanLicenseCode}`, options)
      .catch(this.baseApiService.handleError)
      .map((response: any) => {
        if (response.ok && response._body) {
          let data = response.json() as PlanModel;
          return data;
        } else {
          return null;
        }
      });
    return response;
  }

  getPlanByName(planName: string): Observable<PlanModel> {
    let response = new Observable<PlanModel>(observable => {
      this.getPlans()
        .catch(error => {
          return Observable.throw(error);
        })
        .subscribe((plans: PlanModel[]) => {
          let requestedPlan: PlanModel = {};
          plans.forEach(plan => {
            planName = planName && planName.toLocaleLowerCase();
            let name = plan.Name && plan.Name.toLowerCase();
            if (planName == name) {
              requestedPlan = plan;
              return;
            }
          });
          observable.next(requestedPlan);
        })
    })
    return response;
  }
}
