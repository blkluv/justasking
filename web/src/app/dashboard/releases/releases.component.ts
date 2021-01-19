import { Component, OnInit } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ReleaseService } from '../../core/services/release.service';
import { NotificationsService } from '../../core/services/notifications.service';
import { FeatureRequestService } from '../../core/services/feature-request.service';

import { ReleaseModel } from '../../core/models/release.model';
import { FeatureRequest } from '../../core/models/feature-request.model';

@Component({
  selector: 'app-releases',
  templateUrl: './releases.component.html',
  styleUrls: ['./releases.component.scss']
})
export class ReleasesComponent implements OnInit {

  releases:ReleaseModel[];
  featureRequest:string;
  processingFeatureRequestSubmit:boolean;

  constructor(
    private releaseService:ReleaseService,
    private notificationsService:NotificationsService,
    private featureRequestService:FeatureRequestService,
  ) { }

  submitFeatureRequest(){
    this.processingFeatureRequestSubmit = true;
    let featureRequest: FeatureRequest = {
      FeatureRequest: this.featureRequest
    };
    this.featureRequestService.submitFeatureRequest(featureRequest)
    .catch(error=>{
      this.notificationsService.openSnackBarError("Oh boy, there was an error :( please try again later.");
      this.processingFeatureRequestSubmit = false;
      return Observable.throw(error);
    })
    .subscribe((releases:ReleaseModel[])=>{
      this.featureRequest = "";
      this.processingFeatureRequestSubmit = false;
      this.notificationsService.openSnackBar("Thank you for making justasking.io better!", 3500);
    });
  }  

  ngOnInit() {
    this.releaseService.getReleases()
    .catch(error=>{
      return Observable.throw(error);
    })
    .subscribe((releases:ReleaseModel[])=>{
      this.releases = releases;
    });
  }
}
