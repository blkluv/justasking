import { Injectable } from '@angular/core';

import { FeatureSectionModel } from '../models/feature-section.model';
import { MembershipDetailsModel } from '../models/membership-details.model';

@Injectable()
export class FeatureService {

  constructor() { }

  getBasicFeatureSections(): FeatureSectionModel[]{
    let featureSections = [
      {
        Features: [
          { Name: "Unlimited polls", Active: true },
          { Name: "Unlimited responses", Active: true },
          { Name: "Up to 1 active poll", Active: true },
        ],
      },
      {
        Features: [
          { Name: "SMS enabled polls", Active: false },
          { Name: "Custom poll URL", Active: false },
          { Name: "Account users", Active: false },
          { Name: "Email Support", Active: false },
          // { Name: "Branding - coming soon", Active: false }
        ],
      }
    ];
    return featureSections;
  }

  getPremiumFeatureSections(): FeatureSectionModel[]{
    let featureSections = [
      {
        Features: [
          { Name: "Unlimited polls", Active: true },
          { Name: "Unlimited responses", Active: true },
          { Name: `Up to 5 active polls`, Active: true },
        ],
      },
      {
        Features: [
          { Name: "SMS enabled polls", Active: true },
          { Name: "Custom poll URL", Active: true },
          { Name: "5 Account users", Active: true },
          { Name: "Email Support", Active: true },
          // { Name: "Branding - coming soon", Active: false },
        ],
      },
    ];
    return featureSections;
  }

  getFeatureSectionsFromMembershipDetails(details:MembershipDetailsModel) : FeatureSectionModel[]{
    let featureSections = [
      {
        Features: [
          { Name: "Unlimited polls", Active: true },
          { Name: "Unlimited responses", Active: true },
          { Name: `Up to ${details.ActiveBoxesLimit} active polls`, Active: details.ActiveBoxesLimit > 0 },
        ],
      },
      {
        Features: [
          { Name: "SMS enabled polls", Active: details.Sms },
          { Name: "Custom poll URL", Active: details.CustomCode },
          { Name: `${details.Delegates == 0?'No account users': details.Delegates == 1? '1 Account user': details.Delegates +' Account users'}`, Active: details.Delegates  > 0 },
          { Name: "Email Support", Active: details.Support },
          // { Name: "Branding - coming soon", Active: false },
        ],
      },
    ];
    return featureSections;    
  }

}
