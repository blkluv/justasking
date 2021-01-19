import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { VotesBoxEntryComponent } from './votes-box-entry/votes-box-entry.component';
import { VotesBoxPresentationComponent } from './votes-box-presentation/votes-box-presentation.component';

const routes: Routes = [
    { path: 'presentation', component: VotesBoxPresentationComponent},
    { path: '', component: VotesBoxEntryComponent },
    { path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class VotesBoxRoutingModule {}