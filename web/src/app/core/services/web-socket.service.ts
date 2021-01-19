import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';

import { $WebSocket, WebSocketSendMode, WebSocketConfig } from 'angular2-websocket/angular2-websocket';

@Injectable()
export class WebSocketService {

  activeWebSockets: { ChannelName?: string, $WebSocket?: $WebSocket }[];

  constructor() {
    this.activeWebSockets = [];
  }

  connectToBox(code: string): $WebSocket {
    let ws: $WebSocket;
    let connectedWebSocket = this.isConnectedToChannel(code);
    if (connectedWebSocket) {
      ws = connectedWebSocket;
    } else {
      let sanitizedChannelName = this.sanitizeChannelName(code);
      let wsUrl = `${environment.realTimeBaseDomain}/hubs/box/${sanitizedChannelName}`;
      ws = this.connectToSocket(wsUrl);
      this.activeWebSockets.push({ ChannelName: code, $WebSocket: ws });
    }
    return ws;
  }

  connectToBoxDashboard(code: string): $WebSocket {
    let ws: $WebSocket;
    let connectedWebSocket = this.isConnectedToChannel(code);
    if (connectedWebSocket) {
      ws = connectedWebSocket;
    } else {
      let sanitizedChannelName = this.sanitizeChannelName(code);
      let wsUrl = `${environment.realTimeBaseDomain}/hubs/box/dashboard/${sanitizedChannelName}`;
      ws = this.connectToSocket(wsUrl);
      this.activeWebSockets.push({ ChannelName: code, $WebSocket: ws });
    }
    return ws;
  }

  connectToAccountChannel(accountId: string): $WebSocket {
    let ws: $WebSocket;
    let connectedWebSocket = this.isConnectedToChannel(accountId);
    if (connectedWebSocket) {
      ws = connectedWebSocket;
    } else {
      let sanitizedChannelName = this.sanitizeChannelName(accountId);
      let wsUrl = `${environment.realTimeBaseDomain}/hubs/account/${sanitizedChannelName}`;
      ws = this.connectToSocket(wsUrl);
      this.activeWebSockets.push({ ChannelName: accountId, $WebSocket: ws });
    }
    return ws;
  }

  disconnectWebSockets() {
    this.activeWebSockets.forEach((webSocket) => {
      webSocket.$WebSocket.close();
    });
    this.activeWebSockets = [];
  }

  private sanitizeChannelName(channelName: string): string {
    let sanitizedChannelName = channelName && channelName.trim().toLowerCase().replace(" ", "");
    return sanitizedChannelName;
  }

  private isConnectedToChannel(channelName: string): $WebSocket {
    let connectedWebsocket = null;
    this.activeWebSockets.forEach(webSocket => {
      if (webSocket.ChannelName === channelName) {
        connectedWebsocket = webSocket.$WebSocket;
        return;
      }
    });
    return connectedWebsocket;
  }

  private connectToSocket(socketUrl: string): $WebSocket {
    let ws = new $WebSocket(
      `${environment.websocketProtocol}://${socketUrl}`,
      null,
      { reconnectIfNotNormalClose: true } as WebSocketConfig
    );
    ws.setSend4Mode(WebSocketSendMode.Direct);
    return ws;
  }
}
