import { Observable } from 'rxjs/Observable';
import { FeatureSectionModel } from '../models/feature-section.model';

import { PlanModel } from './plan.model';

export class PricingPlanCardModel {
  PlanType?: string;
  IsPremium?: boolean;
  IsCustom?: boolean;
  CustomPlan?: PlanModel;
  PremiumDurations?: { Duration:string; Price:number; Selected?: boolean; }[];
  SelectetdDuration?: { Duration:string; Price:number; Selected?: boolean; };
  FeatureSections?: FeatureSectionModel[];
  ConfirmText?: string;
  ConfirmAction?: (obj?:any)=> Observable<any>;
}
