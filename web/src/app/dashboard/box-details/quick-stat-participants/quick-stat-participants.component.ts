import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-quick-stat-participants',
  templateUrl: './quick-stat-participants.component.html',
  styleUrls: ['./quick-stat-participants.component.scss']
})
export class QuickStatParticipantsComponent implements OnInit {

  @Input('count') count : number;

  constructor() { }

  ngOnInit() {
  }

}
