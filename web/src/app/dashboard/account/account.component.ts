import { Component, OnInit } from '@angular/core';
import { DataSource } from '@angular/cdk/collections';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/of';
import 'rxjs/add/observable/throw';

import { UserService } from '../../core/services/user.service';

import { UserModel } from '../../core/models/user.model';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnInit {


  constructor(
    public userService: UserService 
  ) {
  }

  ngOnInit() {
  }
}
