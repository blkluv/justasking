import { Injectable } from '@angular/core';
import { MatSnackBar } from '@angular/material';

@Injectable()
export class NotificationsService {

  constructor(
    private snackBar: MatSnackBar    
  ) { }

  openSnackBar(message:string, duration:number){
    let action = "DISMISS";
    this.snackBar.open(message, action, {
      duration: duration,
    });
  }
  
  openSnackBarDefault(message:string){
    let action = "DISMISS";
    this.snackBar.open(message, action, {
      duration: 5000,
    });
  }

  openSnackBarError(message:string){
    let action = "DISMISS";
    this.snackBar.open(`ERROR: ${message}`, action, {
      duration: 5000,
    });
  }
}
