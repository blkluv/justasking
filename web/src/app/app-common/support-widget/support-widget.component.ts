import { Component, OnInit, ViewEncapsulation, ViewChild } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { MatMenuTrigger } from '@angular/material';

import { SupportIssue } from '../../core/models/support-issue';

import { UserService } from '../../core/services/user.service';
import { SupportService } from '../../core/services/support.service';
import { error } from 'selenium-webdriver';

@Component({
  selector: 'app-support-widget',
  templateUrl: './support-widget.component.html',
  styleUrls: ['./support-widget.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class SupportWidgetComponent implements OnInit {

  @ViewChild('supportWidgetTrigger') supportWidgetTrigger:MatMenuTrigger; 
  issue:string;
  processing:boolean;
  submitted:boolean;

  constructor(
    private supportService:SupportService,
    public userService:UserService
  ) { 
    this.processing = false;
    this.submitted = false;
  }

  handleClick(event) {
    event.stopPropagation();
  }

  closeMenu(){
    this.clearForm();
    this.supportWidgetTrigger.closeMenu();
  }

  clearForm(){
    this.processing = false;
    this.submitted = false;
    this.issue = "";
  }

  submitSupportIssue(){
    this.processing = true;
    let supportIssue:SupportIssue = {
      UserAgent:navigator.userAgent,
      Issue:this.issue
    };
    this.supportService.submitIssueToSupport(supportIssue)
    .catch(error=>{
      this.processing = false;
      return Observable.throw(error);
    })
    .subscribe(()=>{
      this.processing = false;
      this.submitted = true;
      setTimeout(()=>{   
        this.closeMenu();
      },4000);
    });
  }

  ngOnInit() {  
  }

}
