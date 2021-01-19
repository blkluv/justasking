import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-no-access',
  templateUrl: './no-access.component.html',
  styleUrls: ['./no-access.component.scss']
})
export class NoAccessComponent implements OnInit {

  @Input("pageTitle") pageTitle: string;

  constructor() { }

  ngOnInit() {
  }

}
