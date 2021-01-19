import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { BoxesComponent } from './boxes.component';

const routes: Routes = [
    { path: 'details/:code', loadChildren: 'app/dashboard/box-details/box-details.module#BoxDetailsModule'},  
    { path: '', component: BoxesComponent },
    { path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class BoxesRoutingModule {}