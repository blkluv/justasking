import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { MatDialog } from '@angular/material';

import { Observable } from 'rxjs/Observable';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { AccountUpgradedDialogComponent } from './account-upgraded-dialog/account-upgraded-dialog.component';

import { BoxService } from '../../core/services/box.service';
import { UserService } from '../../core/services/user.service';
import { WebSocketService } from '../../core/services/web-socket.service';
import { NotificationsService } from '../../core/services/notifications.service';

import { WebsocketMessageModel } from '../../core/models/websocket-message.model';
import { BaseBoxModel } from '../../core/models/base-box.model';
import { UserModel } from '../../core/models/user.model';

@Component({
  selector: 'app-boxes',
  templateUrl: './boxes.component.html',
  styleUrls: ['./boxes.component.scss']
})
export class BoxesComponent implements OnInit {

  boxes: any[];
  loadingBoxes: boolean;
  accountWebSocket: $WebSocket;

  constructor(
    private dialog: MatDialog,
    private userService: UserService,
    private boxService: BoxService,
    private route: ActivatedRoute,
    private webSocketService: WebSocketService,
    private notificationsService: NotificationsService
  ) {
    this.boxes = [];
  }

  private onAccountWebSocketMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    if (response.MessageType === "OpenBox") {
      let box = JSON.parse(response.MessageData) as BaseBoxModel;
      this.boxes.forEach((existingBox: BaseBoxModel) => {
        if (existingBox.Code === box.Code) {
          existingBox.IsLive = box.IsLive;
          return;
        }
      });
    } else if (response.MessageType === "CloseBox") {
      let box = JSON.parse(response.MessageData) as BaseBoxModel;
      this.boxes.forEach((existingBox: BaseBoxModel) => {
        if (existingBox.Code === box.Code) {
          existingBox.IsLive = box.IsLive;
          return;
        }
      });
    } else if (response.MessageType === "NewBox") {
      if (this.userService.user.RolePermissions.SeeAllBoxes) {
        let box = JSON.parse(response.MessageData) as BaseBoxModel;
        this.boxes.unshift(box);
        this.notificationsService.openSnackBarDefault("A new poll has been created!");
      }
    }
  }

  ngOnInit() {
    this.loadingBoxes = true;

    let accountUpgraded = this.route.snapshot.queryParams['upgraded'] == 'true';
    if (accountUpgraded) {
      setTimeout(() => {
        let dialogRef = this.dialog.open(AccountUpgradedDialogComponent);
        dialogRef.afterClosed().subscribe(result => {
        });
      })
    }

    this.userService.getUser()
      .catch(error => {
        this.loadingBoxes = false;
        return Observable.throw(error);
      })
      .subscribe(() => {
        if (this.userService.user) {
          this.accountWebSocket = this.webSocketService.connectToAccountChannel(this.userService.user.Account.Id);
          this.accountWebSocket.onMessage(this.onAccountWebSocketMessage, { autoApply: false });
          this.boxService.getAllBoxes(this.userService.user)
            .catch(error => {
              this.loadingBoxes = false;
              return Observable.throw(error);
            })
            .subscribe((data: BaseBoxModel[]) => {
              this.boxes = data;
              this.loadingBoxes = false;
            });
        }
      });
  }
}