import { Component, OnInit, Input } from '@angular/core';
import { MatDialog } from '@angular/material';
import { Router } from '@angular/router';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { AccountUpgradePromptDialogComponent } from '../../../app-common/account-upgrade-prompt-dialog/account-upgrade-prompt-dialog.component';
import { ConfirmationDialogComponent } from '../../../app-common/confirmation-dialog/confirmation-dialog.component';

import { ProcessingService } from '../../../core/services/processing.service';
import { BoxService } from '../../../core/services/box.service';
import { NotificationsService } from '../../../core/services/notifications.service';

import { BaseBoxModel } from '../../../core/models/base-box.model';
import { WebsocketMessageModel } from '../../../core/models/websocket-message.model';

@Component({
  selector: 'app-quick-stat-box',
  templateUrl: './quick-stat-box.component.html',
  styleUrls: ['./quick-stat-box.component.scss'],
})
export class QuickStatBoxComponent implements OnInit {

  @Input('box') box: BaseBoxModel;
  @Input('boxWebSocket') boxWebSocket: $WebSocket;
  @Input('accountWebSocket') accountWebSocket: $WebSocket;

  boxChannel: $WebSocket;
  accountChannel: $WebSocket;
  pollUrl: string;
  advancedActionsAreVisible: boolean;

  constructor(
    private dialog: MatDialog,
    private notificationsService: NotificationsService,
    private boxService: BoxService,
    private processingService: ProcessingService,
    private router: Router
  ) {
  }

  linkCopied() {
    this.notificationsService.openSnackBarDefault(`Copied poll link to clipboard.`);
  }

  openBox() {
    let baseBoxId = this.box.ID;
    this.processingService.enableProcessingAnimation();
    this.boxService.openBox(baseBoxId)
      .catch(response => {
        if (response.status === 403) {
          this.processingService.disableProcessingAnimation();
          let dialogRef = this.dialog.open(AccountUpgradePromptDialogComponent);
          dialogRef.afterClosed().subscribe(result => {
          });
        } else {
          this.processingService.disableProcessingAnimation();
          this.notificationsService.openSnackBarError("Could not open box.");
        }
        return Observable.throw(response);
      })
      .subscribe(response => {
        this.processingService.disableProcessingAnimation();
        this.box.IsLive = true;
        this.box.PhoneNumber = response.PhoneNumber;
        let box = {
          Code: this.box.Code,
          IsLive: this.box.IsLive,
          PhoneNumber: this.box.PhoneNumber
        }
        let boxString = JSON.stringify(box);
        let websocketMessage: WebsocketMessageModel = { MessageType: "OpenBox", MessageData: boxString }
        this.boxChannel.send(websocketMessage);
        this.accountChannel.send(websocketMessage);
        this.notificationsService.openSnackBarDefault("Box is open for entries.");
      });
  }

  closeBox() {
    let baseBoxId = this.box.ID;
    this.processingService.enableProcessingAnimation();
    this.boxService.closeBox(baseBoxId)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        return Observable.throw(error);
      })
      .subscribe(response => {
        this.processingService.disableProcessingAnimation();
        this.box.IsLive = false;
        this.box.PhoneNumber = "";
        let box = {
          Code: this.box.Code,
          IsLive: this.box.IsLive,
          PhoneNumber: this.box.PhoneNumber
        }
        let boxString = JSON.stringify(box);
        let websocketMessage: WebsocketMessageModel = { MessageType: "CloseBox", MessageData: boxString }
        this.boxChannel.send(websocketMessage);
        this.accountChannel.send(websocketMessage);
        this.notificationsService.openSnackBarDefault("Box is no longer accepting entries.");
      });
  }

  confirmDeleteBox() {
    let baseBoxCode = this.box.Code;
    let baseBoxId = this.box.ID;
    let dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        Header: "Delete poll",
        Texts: [
          "Are you sure you want to delete this poll?", 
          `Access code <b>${baseBoxCode}</b> will be released and your data will be lost.`,
        ],
        ConfirmText: "Yes, delete this poll",
        ConfirmAction: () => { return this.deleteBox(baseBoxId); }
      }
    });
  }

  deleteBox(baseBoxId) : Observable<any> {
    let observableResponse = new Observable<any>(observable => {
    this.processingService.enableProcessingAnimation();
    this.boxService.deleteBox(baseBoxId)
      .catch(error => {
        this.processingService.disableProcessingAnimation();
        observable.next(false);
        this.notificationsService.openSnackBarError("Could not delete this poll. Please try again");
        return Observable.throw(error);
      })
      .subscribe(response => {
        this.processingService.disableProcessingAnimation();
        this.notificationsService.openSnackBarDefault("Successfully deleted poll");
        this.router.navigate(["/dashboard/boxes"]);
        observable.next(true);
      });
    });    
    return observableResponse;
  }

  private onQuickStatBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let box = JSON.parse(response.MessageData) as BaseBoxModel;
    if (response.MessageType == "OpenBox") {
      this.box.IsLive = box.IsLive;
      this.box.PhoneNumber = box.PhoneNumber;
      this.notificationsService.openSnackBarDefault("Box is open for entries.");
    }
    if (response.MessageType == "CloseBox") {
      this.box.IsLive = box && box.IsLive;
      this.box.PhoneNumber = box && box.PhoneNumber;
      this.notificationsService.openSnackBarDefault("Box is no longer accepting entries.");
    }
  }

  ngOnInit() {
    this.boxChannel = this.boxWebSocket;
    this.accountChannel = this.accountWebSocket;
    this.boxChannel.onMessage(this.onQuickStatBoxHubMessage, { autoApply: false });
    this.pollUrl = `https://justasking.io/${this.box.Code}`;
  }

}
