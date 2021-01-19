import { ModuleWithProviders } from '@angular/core';
import { NgModule }             from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { EntryBoxComponent } from './entry-box/entry-box.component';
import { PresentationBoxComponent } from './presentation-box/presentation-box.component';

export const routes: Routes = [
    {
        path: 'presentation',
        component: PresentationBoxComponent},
    {
        path: '',
        component: EntryBoxComponent},
    { 
        path: '**', 
        redirectTo: '404', 
        pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BaseBoxRoutingModule {}
