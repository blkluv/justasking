import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { WordCloudEntryComponent } from './word-cloud-entry/word-cloud-entry.component';
import { WordCloudPresentationComponent } from './word-cloud-presentation/word-cloud-presentation.component';

const routes: Routes = [
    { path: 'presentation', component: WordCloudPresentationComponent },
    { path: '', component: WordCloudEntryComponent },
    { path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class WordCloudRoutingModule {}