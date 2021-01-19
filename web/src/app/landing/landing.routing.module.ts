import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { LandingComponent } from './landing.component';
import { GetstartedComponent } from './getstarted/getstarted.component';
import { PricingComponent } from './pricing/pricing.component';

const routes: Routes = [
{ path: '', component: LandingComponent, 
    children:[
        { path: 'pricing', component: PricingComponent },
        { path: '', component: GetstartedComponent }
    ]
},
{ path: '**', redirectTo: '', pathMatch: 'full' },
];

@NgModule({
imports: [RouterModule.forChild(routes)],
exports: [RouterModule]
})
export class LandingRoutingModule {}