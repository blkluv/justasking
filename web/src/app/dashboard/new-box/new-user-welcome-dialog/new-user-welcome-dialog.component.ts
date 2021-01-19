import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { MatDialog, MatDialogRef } from '@angular/material';
import { Router } from '@angular/router';
import { Location } from '@angular/common';

import { UserService } from '../../../core/services/user.service';

@Component({
  selector: 'app-new-user-welcome-dialog',
  templateUrl: './new-user-welcome-dialog.component.html',
  styleUrls: ['./new-user-welcome-dialog.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class NewUserWelcomeDialogComponent implements OnInit {
  
  constructor(
    public dialogRef: MatDialogRef<NewUserWelcomeDialogComponent>,
    private router: Router,
    private location: Location,
    public userService: UserService
  ) { }
  
  closeDialog(){
    this.location.go("/dashboard/new");
    this.dialogRef.close();
  }

  ngOnInit() {
  }

}
