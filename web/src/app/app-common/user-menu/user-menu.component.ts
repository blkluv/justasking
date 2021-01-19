import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';

import { BaseApiService } from '../../core/services/base-api.service';
import { UserService } from '../../core/services/user.service';
import { IdpAuthenticationService } from '../../core/services/idp-authentication.service';
import { GoogleAnalyticsService } from '../../core/services/google-analytics.service';
import { ImageService } from '../../core/services/image.service';

import { UserModel } from '../../core/models/user.model';

@Component({
  selector: 'app-user-menu',
  templateUrl: './user-menu.component.html',
  styleUrls: ['./user-menu.component.scss']
})
export class UserMenuComponent implements OnInit {

  @Input('user') user: UserModel;
  
  constructor(
    private idpAuthenticationService:IdpAuthenticationService,
    private baseApiService:BaseApiService,
    public userService:UserService,
    private googleAnalyticsService:GoogleAnalyticsService,
    private router:Router,
    private imageService:ImageService
  ) { 
  }

  setDefaultAvatar(event){
    this.imageService.setDefaultProfileAvatar(event);
  }

  handleLogout(){
    this.googleAnalyticsService.trackCustomEvent("logout",{'method': 'User menu - Top right'});
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
  }
}
