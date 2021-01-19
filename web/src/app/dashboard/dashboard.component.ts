import { Component, ViewChild, OnInit } from '@angular/core';
import { Router, NavigationEnd } from '@angular/router';
import { MatSidenav } from '@angular/material';
import { MatDialog } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { BaseApiService } from '../core/services/base-api.service';
import { ProcessingService } from '../core/services/processing.service';
import { SidenavService } from '../core/services/sidenav.service';
import { ReleaseService } from '../core/services/release.service';
import { IdpAuthenticationService } from '../core/services/idp-authentication.service';
import { GoogleAnalyticsService } from '../core/services/google-analytics.service';
import { UserService } from '../core/services/user.service';
import { UserModel } from '../core/models/user.model';
import { ReleaseModel } from '../core/models/release.model';

import { TermsAndConditionsDialogComponent } from '../app-common/terms-and-conditions-dialog/terms-and-conditions-dialog.component';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  @ViewChild('dashboardSidenav') public dashboardSidenav: MatSidenav;
  processing: boolean; 
  latestRelease: ReleaseModel;

  constructor(
    private router:Router,
    private dialog: MatDialog,
    private baseApiService:BaseApiService,
    private sidenavService:SidenavService,
    public userService:UserService,
    private idpAuthenticationService:IdpAuthenticationService,
    public processingService:ProcessingService,
    private releaseService:ReleaseService,
    private googleAnalyticsService:GoogleAnalyticsService
  ) {  
    this.processing = false; 
    router.events.subscribe( (event) => {
      if(event instanceof NavigationEnd) {
        this.sidenavService.close(); 
      } 
    });
    this.latestRelease = {};
  } 
  
  showTermsAndConditions(){
    let dialogRef = this.dialog.open(TermsAndConditionsDialogComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }
  
  toggleSidenav(){
    this.sidenavService.toggle(); 
  }
  
  handleLogout(){
    this.googleAnalyticsService.trackCustomEvent("logout",{'method': 'Leftnav logout button'});
    this.idpAuthenticationService.disconnectGoogleAccount();
    this.baseApiService.removeToken();
    this.userService.resetUser();
    this.router.navigate(["/"]);
  }

  logout(){
    this.userService.logout()
    .catch(error=>{
      this.handleLogout();
      return Observable.throw(error);
    })
    .subscribe(response=>{
      this.handleLogout();
    });
  } 

  ngOnInit() {
    this.sidenavService.setSidenav(this.dashboardSidenav);
    this.userService.getUser()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe();
    this.releaseService.getLatestRelease()
    .catch(error=>{
      return Observable.throw(error);
    })
    .subscribe((release:ReleaseModel)=>{
      this.latestRelease = release;
    });

  }
}
