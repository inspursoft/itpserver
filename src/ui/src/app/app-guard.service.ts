import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, CanActivateChild, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { KeycloakAuthGuard, KeycloakService } from 'keycloak-angular';
import { SharedService } from './shared/shared.service';
import { Observable, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

// @Injectable({providedIn: 'root'})
// export class AppAuthGuardService implements CanActivate, CanActivateChild {
//   constructor(private sharedService: SharedService) {
//   }
//
//   canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):
//     Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
//     return this.sharedService.getVmList().pipe(map(() => true), catchError((res) => {
//       console.log(res);
//       return of(false);
//     }));
//   }
//
//   canActivateChild(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):
//     Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
//     return this.canActivate(route, state);
//   }
// }


@Injectable({providedIn: 'root'})
export class AppAuthGuardService extends KeycloakAuthGuard {
  constructor(protected router: Router, protected keycloakAngular: KeycloakService) {
    super(router, keycloakAngular);
  }

  isAccessAllowed(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean> {
    console.log('isAccessAllowed');
    return new Promise((resolve, reject) => {
      if (!this.authenticated) {
        this.keycloakAngular.login({redirectUri: 'http://10.164.17.1/compatibility/installation'}).then(() => {
          this.keycloakAngular.getToken().then(res => console.log(res));
        });
        return;
      }

      const requiredRoles = route.data.roles;
      if (!requiredRoles || requiredRoles.length === 0) {
        return resolve(true);
      } else {
        if (!this.roles || this.roles.length === 0) {
          resolve(false);
        }
        let granted = false;
        for (const requiredRole of requiredRoles) {
          if (this.roles.indexOf(requiredRole) > -1) {
            granted = true;
            break;
          }
        }
        resolve(granted);
      }
    });
  }
}
