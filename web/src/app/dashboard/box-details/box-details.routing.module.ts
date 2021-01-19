import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { BoxDetailsComponent } from './box-details.component';

const routes: Routes = [
{ path: '', component: BoxDetailsComponent },
{ path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class BoxDetailsRoutingModule {}