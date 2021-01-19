import { Component, OnInit } from '@angular/core';

import { ProcessingService } from '../core/services/processing.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(
    public processingService: ProcessingService,
  ) { }

  ngOnInit() {
  }

}
