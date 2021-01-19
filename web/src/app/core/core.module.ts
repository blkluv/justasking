import { ModuleWithProviders, NgModule,Optional, SkipSelf } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CORE_PROVIDERS } from './core.providers';

@NgModule({
  imports: [CommonModule],
  providers: [CORE_PROVIDERS]
})
export class CoreModule { 

  constructor (@Optional() @SkipSelf() parentModule: CoreModule) {
    if (parentModule) {
      throw new Error(
        'CoreModule is already loaded. Import it in the AppModule only!');
    }
  }

  static forRoot(): ModuleWithProviders {
    return {
      ngModule: CoreModule, 
      providers: []
    };
  }
}
