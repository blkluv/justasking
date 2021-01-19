import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { MatDialog, MatDialogRef } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { Location } from '@angular/common';

import { UserService } from '../../../core/services/user.service';
import { FeatureService } from '../../../core/services/feature.service';

import { FeatureSectionModel } from '../../../core/models/feature-section.model';

@Component({
  selector: 'app-account-upgraded-dialog',
  templateUrl: './account-upgraded-dialog.component.html',
  styleUrls: ['./account-upgraded-dialog.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class AccountUpgradedDialogComponent implements OnInit {

  featureSections: FeatureSectionModel[];

  constructor(
    public dialogRef: MatDialogRef<AccountUpgradedDialogComponent>,
    private userService: UserService,
    private featureService: FeatureService,
    private location: Location
  ) { }

  closeDialog(){
    this.dialogRef.close();
  }

  ngOnInit() {
    this.dialogRef.beforeClose()
    .subscribe(()=>{
      this.location.go("/dashboard/boxes");
    });
    this.userService.getUser()
    .catch(error=>{
      return Observable.throw(error);
    })
    .subscribe((user)=>{
      this.featureSections = this.featureService.getFeatureSectionsFromMembershipDetails(user.MembershipDetails);
    });
  }

}
