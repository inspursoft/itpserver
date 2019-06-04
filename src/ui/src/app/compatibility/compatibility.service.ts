import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Installation, Package, Vm } from './compatibility.type';
import { map } from 'rxjs/operators';
import { plainToClass } from 'class-transformer';

@Injectable({
  providedIn: 'root'
})
export class CompatibilityService {
  constructor(private http: HttpClient) {

  }

  getVmList(): Observable<Array<Vm>> {
    return this.http.get<Array<Vm>>(`/v1/vms`, {
      headers: {
        accept: 'application/json'
      }
    }).pipe(map(res => plainToClass(Vm, res)));
  }

  deleteVm(vmName: string): Observable<any> {
    return this.http.delete(`/v1/vms/${vmName}`, {
      params: {vm_name: vmName}
    });
  }

  deletePackage(packageName, packageTag: string): Observable<any> {
    return this.http.delete(`/v1/packages`, {
      params: {
        package_name: packageName,
        package_tag: packageTag
      }
    });
  }

  getPackageList(): Observable<Array<Package>> {
    return this.http.get<Array<Package>>(`/v1/packages`)
      .pipe(map(res => plainToClass(Package, res)));
  }

  getInstallationList(vmName: string): Observable<Array<Installation>> {
    return this.http.get<Array<Installation>>(`/v1/installations`, {
      params: {
        vm_name: vmName
      }
    }).pipe(map(res => plainToClass(Installation, res)));
  }

  createInstallation(vmName: string, packageName, packageTag: string): Observable<any> {
    return this.http.post(`/v1/installations/${vmName}`, {
      package_name: packageName,
      package_tag: packageTag
    }, {responseType: 'text'});
  }

  deleteInstallation(vmName: string, installation: Installation): Observable<any> {
    return this.http.delete(`/v1/installations/${vmName}`, {
      params: {
        pkg_name: installation.name,
        pkg_tag: installation.tag
      }
    });
  }
}
