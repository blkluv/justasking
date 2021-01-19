import { NgModule } from '@angular/core';

import { AppCommonModule } from '../app-common/app-common.module';

import { VotesBoxEntryComponent } from './votes-box-entry/votes-box-entry.component';
import { VotesBoxPresentationComponent } from './votes-box-presentation/votes-box-presentation.component';
import { VotesBoxDetailsComponent } from './votes-box-details/votes-box-details.component';

@NgModule({
  imports: [
    AppCommonModule
  ],
  exports: [
    VotesBoxEntryComponent, 
    VotesBoxPresentationComponent,
    VotesBoxDetailsComponent
  ],
  declarations: [
    VotesBoxEntryComponent, 
    VotesBoxPresentationComponent,
    VotesBoxDetailsComponent
  ]
})
export class VotesBoxModule { }
