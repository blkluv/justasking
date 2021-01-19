import { ModuleWithProviders } from '@angular/core';
import { NgModule }             from '@angular/core';
import { Routes, RouterModule, PreloadAllModules,  ExtraOptions } from '@angular/router';

import { AuthGuard } from './core/security/auth.guard';

import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { InternalServerErrorComponent } from './internal-server-error/internal-server-error.component';

export const routes: Routes = [
  { 
    path: 'dashboard', 
    loadChildren: 'app/dashboard/dashboard.module#DashboardModule',
    canActivate: [AuthGuard]
  },
  { 
    path: 'getstarted', 
    loadChildren: 'app/landing/landing.module#LandingModule'}, 
  { 
    path: 'login', 
    loadChildren: 'app/login/login.module#LoginModule'
  }, 
  { 
    path: '404',
    component: PageNotFoundComponent}, 
  { 
    path: '500', 
    component: InternalServerErrorComponent},  
  { 
    path: ':code', 
    loadChildren: 'app/base-box/base-box.module#BaseBoxModule'},  
  { 
    path: '', 
    redirectTo: 'getstarted', 
    pathMatch: 'full'},  
  { 
    path: '**', 
    redirectTo: '', 
    pathMatch: 'full' }, 
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {
    onSameUrlNavigation: 'reload'
  })],
  exports: [RouterModule]
})
export class AppRoutingModule {}
