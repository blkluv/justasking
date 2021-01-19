import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { BaseBoxModel } from '../../../core/models/base-box.model';
import { UserModel } from '../../../core/models/user.model';
import { UserService } from '../../../core/services/user.service';
import { NotificationsService } from '../../../core/services/notifications.service';

declare var window: any;
declare var html2canvas: any;

@Component({
  selector: 'app-presentation-topnav',
  templateUrl: './presentation-topnav.component.html',
  styleUrls: ['./presentation-topnav.component.scss']
})
export class PresentationTopnavComponent implements OnInit {

  @Input('box') baseBox: BaseBoxModel;
  @Input('hidePhone') hidePhone: boolean;
  boxUrl:string;

  constructor(
    public notificationsService: NotificationsService,
    public userService: UserService,
    private router: Router
  ) {
  }

  download(code){
    setTimeout(function(){
      html2canvas(document.body, {
          onrendered: function(canvas) {
              var newImg = document.createElement("img");
              newImg.src = canvas.toDataURL();
              var newAnchor = document.createElement("a");
              newAnchor.href = newImg.src;
              newAnchor.download = `${code}_screenshot.png`;
              newAnchor.id="screenshot";
              document.body.appendChild(newAnchor);
              newAnchor.click();
          }
      });
    },200);
  }

  copyToClipboard(){
    this.notificationsService.openSnackBarDefault("Link copied to clipboard!");
  }

  ngOnInit() {
    this.boxUrl = document.location.href;
    let userIsLoggedIn = this.userService.userIsLoggedIn();
    if(userIsLoggedIn){
      this.userService.getUser()
        .catch(error=>{ 
            return Observable.throw(error);
        }).subscribe();
    } 
  }

}
