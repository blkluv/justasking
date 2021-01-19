import { Injectable } from '@angular/core';
import { Router, CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private router: Router) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    let canActivate: boolean = false;
    if (localStorage.getItem('zxcv')) {
      canActivate = true;
    } else {
      this.router.navigate(['/login'], { queryParams: { returnUrl: state.url } });
      canActivate = false;
    } 
    return canActivate;
  }
}