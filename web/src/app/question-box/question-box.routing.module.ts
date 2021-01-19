import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { QuestionBoxPresentationComponent } from './question-box-presentation/question-box-presentation.component';
import { QuestionBoxEntryComponent } from './question-box-entry/question-box-entry.component';

const routes: Routes = [
    { path: 'presentation', component: QuestionBoxPresentationComponent },
    { path: '', component: QuestionBoxEntryComponent },
    { path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class QuestionBoxRoutingModule {}