import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-quick-stat-entries',
  templateUrl: './quick-stat-entries.component.html',
  styleUrls: ['./quick-stat-entries.component.scss']
})
export class QuickStatEntriesComponent implements OnInit {

  @Input('count') count : number;
  @Input('headerText') headerText : string;
  
  constructor() { }

  ngOnInit() {
    if(this.headerText == null || typeof(this.headerText) == 'undefined'){
      this.headerText = "Entries count";
    }
  }

}
