import { BrowserModule } from '@angular/platform-browser';
import { APP_INITIALIZER, NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SharedModule } from './shared/shared.module';
import { MainContentComponent } from './main-content/main-content.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { IconDefinition } from '@ant-design/icons-angular';
import { AccountBookFill, AlertFill, AlertOutline } from '@ant-design/icons-angular/icons';
import { NZ_ICONS } from 'ng-zorro-antd';
import { KeycloakAngularModule, KeycloakService } from 'keycloak-angular';
import { initializer } from './app-init';

const icons: IconDefinition[] = [AccountBookFill, AlertOutline, AlertFill];

@NgModule({
  declarations: [
    AppComponent,
    MainContentComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    KeycloakAngularModule,
    AppRoutingModule,
    SharedModule
  ],
  providers: [
    {
      provide: APP_INITIALIZER,
      useFactory: initializer,
      multi: true,
      deps: [KeycloakService]
    },
    {provide: NZ_ICONS, useValue: icons}],
  bootstrap: [AppComponent]
})
export class AppModule {
}
