import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { BaseApiService } from '../../core/services/base-api.service';
import { UserService } from '../../core/services/user.service';
import { AccountService } from '../../core/services/account.service';
import { NotificationsService } from '../../core/services/notifications.service';

import { AccountUserInviteModel } from '../../core/models/account-user-invite.model';
import { UserModel } from '../../core/models/user.model';

@Component({
  selector: 'app-join-account',
  templateUrl: './join-account.component.html',
  styleUrls: ['./join-account.component.scss']
})
export class JoinAccountComponent implements OnInit {

  invitationCode: string;
  accountUserInvitation: AccountUserInviteModel;

  constructor( 
    private router: Router,
    private route: ActivatedRoute,
    private baseApiService: BaseApiService,
    private userService: UserService,
    private accountService: AccountService,
    private notificationsService: NotificationsService
  ) {  
    this.invitationCode = "";
  }

  joinAccount(){
    this.accountService.joinAccount(this.invitationCode)
    .catch(error=>{
      return Observable.throw(error);
    })
    .subscribe(newToken =>{
      this.notificationsService.openSnackBarDefault(`Joined ${this.accountUserInvitation.AccountName}'s account.`)
      this.baseApiService.setToken(newToken);
      this.baseApiService.updatedToken(newToken);
      this.userService.getUserFromApi()
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe(()=>{ 
        this.router.navigate(["dashboard"]);
      })
    });
  }

  ngOnInit() {
    this.route.params.forEach((params: Params) => {
      this.invitationCode = params["inviteKey"];
    });
    this.userService.getUser()
    .catch(error=>{
      return Observable.throw(error);
    })
    .subscribe((user:UserModel)=>{
      this.accountService.getAccountInvitationByKey(this.invitationCode)
      .catch(error=>{
        return Observable.throw(error);
      })
      .subscribe((accountInvite:AccountUserInviteModel)=>{
        let currentUserEmail = user.Email && user.Email.trim().toLowerCase();
        let invitationEmail = accountInvite.Email && accountInvite.Email.trim().toLowerCase();
        if(currentUserEmail === invitationEmail){
          this.accountUserInvitation = accountInvite;
        }
      });
    })
  }
}
