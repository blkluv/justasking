import { Injectable } from '@angular/core';
import { MatSidenav, MatDrawerToggleResult } from '@angular/material';

@Injectable()
export class SidenavService {
  private sidenav: MatSidenav;

  public setSidenav(sidenav: MatSidenav) {
    this.sidenav = sidenav;
  }
 
  public open(): Promise<any> {
    return this.sidenav.open();
  }
 
  public close(): Promise<any> {
    if(this.sidenav){
      return this.sidenav.close();
    }
    return;
  }
 
  public toggle(isOpen?: boolean): Promise<any> {
    return this.sidenav.toggle(isOpen);
  }
 
  public clearSidenav(){
    this.sidenav = null;
  }
}