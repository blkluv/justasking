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
        question: "Do you offer a paid plan?",
        answer: `JustAsking is free for anyone who finds it useful.`,
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
