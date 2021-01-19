import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef } from '@angular/material';

@Component({
  selector: 'app-terms-and-conditions-dialog',
  templateUrl: './terms-and-conditions-dialog.component.html',
  styleUrls: ['./terms-and-conditions-dialog.component.scss']
})
export class TermsAndConditionsDialogComponent implements OnInit {

  constructor(public dialogRef: MatDialogRef<TermsAndConditionsDialogComponent>) { }

  ngOnInit() {
  }

  closeDialog(){
    this.dialogRef.close();
  }

}
