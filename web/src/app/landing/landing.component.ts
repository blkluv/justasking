import { Component, OnInit } from '@angular/core';

import { UserService } from '../core/services/user.service';
import { UserModel } from '../core/models/user.model';

@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.scss'],
  host: {
      '(window:scroll)': 'onScroll($event)'
  }
})
export class LandingComponent implements OnInit {
  
  isScrolled = false;
  currPos: number = 0;
  startPos: number = 0;
  changePos: number = 100;

  faqs: any[];
  userIsLoggedIn: boolean;

  constructor(
    private userService: UserService 
  ) {
    this.userIsLoggedIn = false;
  }
  
  onScroll(evt) {
    this.currPos = (window.pageYOffset || evt.target.scrollTop) - (evt.target.clientTop || 0);
    this.isScrolled = this.currPos >= this.changePos;
}

  ngOnInit() {
    this.userIsLoggedIn = this.userService.userIsLoggedIn(); 
  }

}
