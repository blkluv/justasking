import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router';
import { MatDialog, MatDialogRef } from '@angular/material';

import { UserService } from '../../core/services/user.service';
import { FeatureService } from '../../core/services/feature.service';

import { FeatureSectionModel } from '../../core/models/feature-section.model';

@Component({
  selector: 'app-account-upgrade-prompt-dialog',
  templateUrl: './account-upgrade-prompt-dialog.component.html',
  styleUrls: ['./account-upgrade-prompt-dialog.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class AccountUpgradePromptDialogComponent implements OnInit {

  featureSections: FeatureSectionModel[];

  constructor(
    public dialogRef: MatDialogRef<AccountUpgradePromptDialogComponent>,
    private router: Router,
    private featureService: FeatureService,
    public userService: UserService,
  ) { }
  
  notNow(){
    this.closeDialog();
  }
    
  closeDialog(){
    this.dialogRef.close();
  }

  upgrade(){
    this.router.navigate(["/dashboard/account/upgrade"]).then(()=>{
      this.dialogRef.close();
    });
  }

  ngOnInit() {
    this.featureSections = this.featureService.getPremiumFeatureSections();    
  }

}
