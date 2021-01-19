import { Injectable, RendererFactory2, Renderer2 } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { BaseApiService } from './base-api.service';

import { PlanModel } from '../../core/models/plan.model';
import { UserModel } from '../../core/models/user.model';

import { environment } from '../../../environments/environment';

@Injectable()
export class StripeService {

  private globalListener: any;
  private renderer: Renderer2;

  constructor(
    private http: Http,
    private baseApiService: BaseApiService,
    private rendererFactory: RendererFactory2
  ) {
    this.renderer = rendererFactory.createRenderer(null, null);
  }

  openCheckout(plan: PlanModel, user: UserModel, callback: (response: any) => void) {
    var handler = (<any>window).StripeCheckout.configure({
      key: environment.stripeKey,
      locale: 'auto',
      token: function (response: any) {
        callback(response);
      }
    });

    handler.open({
      amount: plan.Price * 100,
      name: 'justasking.io',
      description: plan.DisplayName,
      image: '/assets/logos/logo-indigo.png',
      email: user.Email,
      panelLabel: "Pay",
      allowRememberMe: false,
      zipCode: true
    });

    this.globalListener = this.renderer.listen('window', 'popstate', () => {
      handler.close();
    });
  }

  updateUsersDefaultCardOnStripe(card: any): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.post(`${baseUrl}card`, card, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

  getStripeData(): Observable<any> {
    let baseUrl = this.baseApiService.getBaseUrl();
    let options = this.baseApiService.getRequestOptions();
    return this.http.get(`${baseUrl}stripedata`, options)
      .catch(this.baseApiService.handleError)
      .map(this.baseApiService.extractData);
  }

}
