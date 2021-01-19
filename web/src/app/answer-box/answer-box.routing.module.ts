import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AnswerBoxEntryComponent } from './answer-box-entry/answer-box-entry.component';
import { AnswerBoxPresentationComponent } from './answer-box-presentation/answer-box-presentation.component';

const routes: Routes = [
    { path: 'presentation', component: AnswerBoxPresentationComponent },
    { path: '', component: AnswerBoxEntryComponent },
    { path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class AnswerBoxRoutingModule {}