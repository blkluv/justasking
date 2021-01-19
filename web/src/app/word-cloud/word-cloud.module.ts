import { NgModule } from '@angular/core';

import { AppCommonModule } from '../app-common/app-common.module';
import { WordCloudRoutingModule } from './word-cloud.routing.module';

import { WordCloudEntryComponent } from './word-cloud-entry/word-cloud-entry.component';
import { WordCloudPresentationComponent } from './word-cloud-presentation/word-cloud-presentation.component';
import { WordCloudDetailsComponent } from './word-cloud-details/word-cloud-details.component';
import { WordCloudComponent } from './word-cloud.component';

@NgModule({
  imports: [
    WordCloudRoutingModule,
    AppCommonModule,
  ],
  exports:[
    WordCloudEntryComponent,
    WordCloudPresentationComponent, 
    WordCloudDetailsComponent,
    WordCloudComponent
  ],
  declarations: [
    WordCloudEntryComponent, 
    WordCloudPresentationComponent, 
    WordCloudDetailsComponent,
    WordCloudComponent
  ]
})
export class WordCloudModule { }
