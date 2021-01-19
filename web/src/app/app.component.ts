import { Component } from '@angular/core';
import { Router, NavigationEnd, NavigationStart } from '@angular/router';

import { ProcessingService } from './core/services/processing.service'
import { WebSocketService } from './core/services/web-socket.service'
import { GoogleAnalyticsService } from './core/services/google-analytics.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(
    private router:Router,
    private processingService:ProcessingService,
    private webSocketService:WebSocketService,
    private googleAnalyticsService: GoogleAnalyticsService
    ) {
      router.events.subscribe( (event) => {
        if(event instanceof NavigationStart) {
          processingService.enableProcessingAnimation();
          webSocketService.disconnectWebSockets();
        }
        if(event instanceof NavigationEnd) {
          window.scrollTo(0, 0);
          processingService.disableProcessingAnimation();
          this.googleAnalyticsService.trackPageView(event.urlAfterRedirects);
      }
    });
  }
}
