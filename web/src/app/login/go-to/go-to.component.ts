import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from "@angular/router";

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { BoxService } from '../../core/services/box.service';
import { ProcessingService } from '../../core/services/processing.service';
import { NotificationsService } from '../../core/services/notifications.service';

@Component({
  selector: 'app-go-to',
  templateUrl: './go-to.component.html',
  styleUrls: ['./go-to.component.scss']
})
export class GoToComponent implements OnInit {

  goToPollByAccessCodeForm: FormGroup;

  constructor(
    private router:Router,
    private boxService:BoxService,
    private processingService:ProcessingService,
    private notificationsService:NotificationsService
  ) { }

  goToPoll(accessCode:string){
    accessCode = accessCode && accessCode.trim();
    if(accessCode){
      this.processingService.enableProcessingAnimation();
      this.boxService.codeIsTaken(accessCode)
      .catch(error=>{
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarError("There was an error verifying that code. Please try again.");
        return Observable.throw(error);
      })
      .subscribe(code =>{
        this.processingService.disableProcessingAnimation();
        if(!code.Exists){
          this.notificationsService.openSnackBarDefault("That poll does not exist, please verify the access code and try again.")
        }else{
          this.router.navigate([`/${accessCode}`]);
        }
      })
    }else{
      this.notificationsService.openSnackBarDefault("Please type an access code and try again.")
    }
  }

  ngOnInit() {
    this.goToPollByAccessCodeForm = new FormGroup({
      'AccessCode': new FormControl('', [
        Validators.required,
        Validators.maxLength(50)
      ])
    });
  }
}
