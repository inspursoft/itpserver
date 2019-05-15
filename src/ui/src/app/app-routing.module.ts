import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {RouteAccount, RouteCompatibility, RouteLogin} from './shared/shared.const';
import {Error404Component} from './shared/error_pages/error404/error404.component';
import {MainContentComponent} from './main-content/main-content.component';

const routes: Routes = [
  {path: RouteAccount, loadChildren: './account/account.module#AccountModule'},
  {
    path: '', component: MainContentComponent, children: [
      {path: RouteCompatibility, loadChildren: './compatibility/compatibility.module#CompatibilityModule'}
    ]
  },
  {path: '**', component: Error404Component}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
