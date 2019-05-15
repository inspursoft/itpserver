import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { VmListComponent } from './vm/vm-list/vm-list.component';
import { InstallationListComponent } from './installation/installation-list/installation-list.component';
import { PackageListComponent } from './package/package-list/package-list.component';
import { SharedModule } from '../shared/shared.module';
import { RouteBaseOs, RouteInstallation, RoutePackage, RouteVm } from '../shared/shared.const';
import { VmDetailComponent } from './vm/vm-detail/vm-detail.component';
import { BaseOsListComponent } from './base-os-list/base-os-list.component';

@NgModule({
  declarations: [
    VmListComponent,
    InstallationListComponent,
    PackageListComponent,
    VmDetailComponent,
    BaseOsListComponent
  ],
  entryComponents: [
    VmDetailComponent
  ],
  imports: [
    SharedModule,
    RouterModule.forChild([
      {path: RouteVm, component: VmListComponent},
      {path: RoutePackage, component: PackageListComponent},
      {path: RouteInstallation, component: InstallationListComponent},
      {path: RouteBaseOs, component: BaseOsListComponent},
    ])
  ]
})
export class CompatibilityModule {
}
