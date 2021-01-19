import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';

declare let gtag: Function;

@Injectable()
export class GoogleAnalyticsService {

  constructor() { 
  } 

  trackPageView(pagePath: string){
    if(gtag && environment.envName == 'production'){ 
      gtag('config', environment.googleAnalyticsTrackingID, {
        'page_location': window.location.href,
        'page_path': pagePath
      });
    }
  }

  trackSignUp(){
    if(gtag && environment.envName == 'production'){ 
    gtag('event', 'conversion', {'send_to': 'AW-852718222/xzfXCK7y_noQjuXNlgM'});
    }
  }

  trackPremiumUpgradeMonth(){
    if(gtag && environment.envName == 'production'){ 
      gtag('event', 'conversion', {'send_to': 'AW-852718222/5y8bCJvhgXsQjuXNlgM','transaction_id': ''});
    }
  }

  trackPremiumUpgradeYear(){
    if(gtag && environment.envName == 'production'){ 
      gtag('event', 'conversion', {'send_to': 'AW-852718222/wd1LCL_hgXsQjuXNlgM','transaction_id': ''});
    }
  }

  trackCustomEvent(eventName:string, customEventObject:any){
    if(gtag && environment.envName == 'production'){ 
      gtag('event', eventName, customEventObject);
    }
  }
}