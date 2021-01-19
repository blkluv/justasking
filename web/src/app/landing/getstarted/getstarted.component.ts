import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef } from '@angular/material';

declare var videojs:any;

@Component({
  selector: 'app-getstarted',
  templateUrl: './getstarted.component.html',
  styleUrls: ['./getstarted.component.scss']
})
export class GetstartedComponent implements OnInit {

  faqs: any[];
  videoOverlayOn: boolean;

  constructor(private dialog: MatDialog) {
  }

  toggleVideoOverlay(){
    this.videoOverlayOn = !this.videoOverlayOn;
  }

  ngOnInit() { 
    this.faqs = [
      {
        question: "Do you offer a free plan?",
        answer: `Yes, with our BASIC plan you can create unlimited number of polls for unlimited number of participants. You will, however be missing out 
        on some cool features like receiving responses through SMS, having multiple polls open at the same time, customizing polls with your branding, 
        sharing control of your polls with others, and more.`,
        expanded: false
      },
      {
        question: "How does billing work?",
        answer: `When you sign up for the PREMIUM plan, you will be billed once. You can choose to keep the plan for one month or one year.`,
        expanded: false
      },
      {
        question: "Which devices are supported?",
        answer: "Any device that has access to the internet or can submit an SMS, is supported.",
        expanded: false
      },
      {
        question: "What if i'm not in the United States?",
        answer: "Not a problem! You can still use our online service to create and interact with polls.",
        expanded: false
      },
    ]
  }

}
