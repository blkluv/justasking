import { Component, Input, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { $WebSocket } from 'angular2-websocket/angular2-websocket';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { BaseBoxModel } from '../../core/models/base-box.model';
import { WebsocketMessageModel } from '../../core/models/websocket-message.model';

import { BoxService } from '../../core/services/box.service';
import { WebSocketService } from '../../core/services/web-socket.service';

@Component({
  selector: 'app-presentation-box',
  templateUrl: './presentation-box.component.html',
  styleUrls: ['./presentation-box.component.scss']
})
export class PresentationBoxComponent implements OnInit {

  @Input('boxPreview') boxPreview: BaseBoxModel;

  box: any;
  ws: $WebSocket;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private webSocketService: WebSocketService,
    private boxService: BoxService
  ) {
    this.box = {}
  }

  private onPresentationBoxHubMessage = (msg: MessageEvent) => {
    let response: WebsocketMessageModel = JSON.parse(msg.data);
    let box = JSON.parse(response.MessageData);
    if (response.MessageType == "OpenBox") {
      this.box.BaseBox.IsLive = true;
      this.box.BaseBox.PhoneNumber = box.PhoneNumber;
    }
    if (response.MessageType == "CloseBox") {
      this.box.BaseBox.IsLive = false;
      this.box.BaseBox.PhoneNumber = box && box.PhoneNumber;
    }
  }

  ngOnInit() {
    if (this.boxPreview) {
      this.box = this.boxPreview;
    } else {
      let code: string;
      this.route.params.forEach((params: Params) => {
        code = params["code"];
      });
      this.ws = this.webSocketService.connectToBox(code);
      this.ws.onMessage(this.onPresentationBoxHubMessage, { autoApply: false });
      this.boxService.getBaseBox(code)
        .catch(error => {
          if (error.status == 500) {
            this.router.navigate(["500"])
          }
          if (error.status == 404) {
            this.router.navigate(["404"])
          }
          return Observable.throw(error);
        })
        .subscribe(baseBox => {
          this.box = { BaseBox: baseBox };
        });
    }
  }
}
