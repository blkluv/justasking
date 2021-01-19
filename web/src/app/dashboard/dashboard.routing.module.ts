import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { DashboardComponent } from './dashboard.component'; 
import { PageNotFoundComponent } from '../page-not-found/page-not-found.component'; 
import { InternalServerErrorComponent } from '../internal-server-error/internal-server-error.component'; 

const routes: Routes = [
{ path: '', component: DashboardComponent, 
    children: [
        { path: 'boxes', loadChildren: 'app/dashboard/boxes/boxes.module#BoxesModule'},  
        { path: 'new', loadChildren: 'app/dashboard/new-box/new-box.module#NewBoxModule'},  
        { path: 'account', loadChildren: 'app/dashboard/account/account.module#AccountModule'},  
        { path: 'feedback', loadChildren: 'app/dashboard/feedback/feedback.module#FeedbackModule'},  
        { path: 'releases', loadChildren: 'app/dashboard/releases/releases.module#ReleasesModule'},  
        { path: '404', component: PageNotFoundComponent,},  
        { path: '500', component: InternalServerErrorComponent,},  
        { path: '', redirectTo: 'boxes'},
    ]
},
{ path: '', component: DashboardComponent },
{ path: '**', redirectTo: '404', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class DashboardRoutingModule {}