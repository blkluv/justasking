import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';
import 'rxjs/add/operator/catch';

import { ReleaseModel } from '../../core/models/release.model';

@Injectable()
export class ReleaseService {

  releases : ReleaseModel[] = [
    {
      Name: "V1.2.0",
      Date: "March. 3, 2018",
      Notes: [
        "As a PREMIUM user, you can now add users to your account.",
        "You may set new users under your account to be presenters, admins or owners.",
        "Redesigned result view for polls. They are now displayed in a horizontal bar graph.",
        "Mobile view enhancements for every Q&A, poll and wordcloud."
      ]
    },
    {
      Name: "V1.0.2",
      Date: "Jan. 12, 2018",
      Notes: [
        "You can now create a justasking account with an email and a password in addition to using Google to sign in.",
        "Share your poll on social media in the poll details page.",
        "Minor bug fixes.",
      ]
    },
    {
      Name: "V1.0.1",
      Date: "Dec. 17, 2017",
      Notes: [
        "Report any issues you're experiencing with justasking.io right from the web application.",
        "A help widget lives on the bottom right corner of the application that helps you submit a summary of the issue.",
      ]
    },
    {
      Name: "V1.0.0",
      Date: "Dec. 29, 2017",
      Notes: [
        "Version 1.0 of justasking.io is finally out.",
        "Create unlimited polls with unlimited responses."
      ]
    },
  ];

  constructor() { }

  getReleases(): Observable<ReleaseModel[]>{
    return new Observable<ReleaseModel[]>(observer => { 
      observer.next(this.releases);
    });
  }

  getLatestRelease(): Observable<ReleaseModel>{
    return new Observable<ReleaseModel>(observer => { 
      observer.next(this.releases[0]);
    });
  }

}
