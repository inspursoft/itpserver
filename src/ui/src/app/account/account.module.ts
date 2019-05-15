import {NgModule} from '@angular/core';
import {RouterModule} from '@angular/router';
import {LoginComponent} from './login/login.component';
import {RouteLogin} from '../shared/shared.const';
import {SharedModule} from '../shared/shared.module';

@NgModule({
  declarations: [
    LoginComponent
  ],
  imports: [
    SharedModule,
    RouterModule.forChild([
      {path: '', redirectTo: `${RouteLogin}`},
      {path: RouteLogin, component: LoginComponent}
    ])
  ]
})
export class AccountModule {

}
