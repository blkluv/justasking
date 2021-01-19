import { Component, OnInit, Inject } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ConfirmationDialogModel } from '../../core/models/confirmation-dialog.model';

@Component({
  selector: 'app-confirmation-dialog',
  templateUrl: './confirmation-dialog.component.html',
  styleUrls: ['./confirmation-dialog.component.scss']
})
export class ConfirmationDialogComponent implements OnInit {

  processing: boolean;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: ConfirmationDialogModel,
    private dialogRef: MatDialogRef<ConfirmationDialogComponent>
  ) {
    this.dialogRef.disableClose = true;
    this.processing = false;
  }

  confirm() {
    this.processing = true;
    this.data.ConfirmAction()
      .catch(error => {
        this.processing = false;
        return Observable.throw(error);
      })
      .subscribe(response => {
        this.processing = false;
        this.closeDialog(response);
      })
  }

  notNow() {
    this.closeDialog();
  }

  closeDialog(data?: any) {
    this.dialogRef.close(data);
  }

  ngOnInit() {
  }

}