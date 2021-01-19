import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { UserModel } from '../../core/models/user.model';
import { BaseBoxModel } from '../../core/models/base-box.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

import { BoxService } from '../../core/services/box.service';
import { UserService } from '../../core/services/user.service';
import { WebSocketService } from '../../core/services/web-socket.service';

@Component({
  selector: 'app-box-details',
  templateUrl: './box-details.component.html',
  styleUrls: ['./box-details.component.scss']
})
export class BoxDetailsComponent implements OnInit {

  boxWebSocket: $WebSocket;
  accountWebSocket: $WebSocket;
  boxClientCount: number;
  box: BaseBoxModel;
  hasAccess: boolean;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private boxService: BoxService,
    private userService: UserService,
    private webSocketService: WebSocketService
  ) {
    this.hasAccess = true;
    this.box = {};
  }

  private onBoxDetailsHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    if (response.MessageType == "DashboardClientCount") {
      let count = JSON.parse(response.MessageData).ClientCount;
      this.boxClientCount = count;
    }
  }

  ngOnInit() {
    let code: string;
    this.route.params.forEach((params: Params) => {
      code = params["code"];
    });
    this.boxWebSocket = this.webSocketService.connectToBoxDashboard(code);
    this.boxWebSocket.onMessage(this.onBoxDetailsHubMessage, { autoApply: false });
    this.userService.getUser()
      .catch(error => {
        return Observable.throw(error);
      })
      .subscribe(() => {
        this.accountWebSocket = this.webSocketService.connectToAccountChannel(this.userService.user.Account.Id);
        this.boxService.getBaseBoxAuth(code)
          .catch(error => {
            if (error.status == 403) {
              this.hasAccess = false;
            }
            if (error.status == 500) {
              this.router.navigate(["/dashboard/500"]);
            }
            if (error.status == 404) {
              this.router.navigate(["/dashboard/404"]);
            }
            return Observable.throw(error);
          })
          .subscribe((baseBox: BaseBoxModel) => {
            if (baseBox.AccountId == this.userService.user.Account.Id) {
              this.box = baseBox;
            } else {
              this.router.navigate(["/dashboard/404"]);
            }
          });
      })
  }
}
