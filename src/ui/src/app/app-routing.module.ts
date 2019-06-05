import { NgModule } from '@angular/core';
import { Routes, RouterModule, Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { RouteCompatibility } from './shared/shared.const';
import { Error404Component } from './shared/error_pages/error404/error404.component';
import { MainContentComponent } from './main-content/main-content.component';
import { AppAuthGuardService } from './app-guard.service';

const routes: Routes = [
  {
    path: '',
    canActivate: [AppAuthGuardService],
    component: MainContentComponent,
    children: [
      {path: RouteCompatibility,
        loadChildren: './compatibility/compatibility.module#CompatibilityModule'}
    ]
  },
  {path: '**', component: Error404Component}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  providers: [AppAuthGuardService],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
