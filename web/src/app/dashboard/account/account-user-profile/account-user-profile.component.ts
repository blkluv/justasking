import { Component, OnInit, Input } from '@angular/core';

import { ImageService } from '../../../core/services/image.service';

import { UserModel } from '../../../core/models/user.model';

@Component({
  selector: 'app-account-user-profile',
  templateUrl: './account-user-profile.component.html',
  styleUrls: ['./account-user-profile.component.scss']
})
export class AccountUserProfileComponent implements OnInit {

  @Input('user') user : UserModel;

  constructor(
    private imageService: ImageService
  ) { }
  
  setDefaultAvatar(event){
    this.imageService.setDefaultProfileAvatar(event);
  }

  ngOnInit() {
  }

}
