import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Error404Component } from './error_pages/error404/error404.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { HttpClientModule } from '@angular/common/http';
import { HttpInterceptorService } from './http-client-interceptor';
import { CreateVmComponent } from './create-vm/create-vm.component';
import { SharedService } from './shared.service';
import { SharedActionService } from './shared.action.service';
import { CreatePackageComponent } from './create-package/create-package.component';
import { LogViewerComponent } from './log-viewer/log-viewer.component';

@NgModule({
  declarations: [
    Error404Component,
    CreateVmComponent,
    CreatePackageComponent,
    LogViewerComponent
  ],
  providers: [
    SharedService,
    SharedActionService,
    HttpInterceptorService
  ],
  entryComponents: [
    CreateVmComponent,
    CreatePackageComponent,
    LogViewerComponent
  ],
  imports: [
    ReactiveFormsModule,
    NgZorroAntdModule,
    FormsModule,
    HttpClientModule,
    CommonModule
  ],
  exports: [
    ReactiveFormsModule,
    NgZorroAntdModule,
    FormsModule,
    HttpClientModule,
    CommonModule
  ]
})
export class SharedModule {
}
