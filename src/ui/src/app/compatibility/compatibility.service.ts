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

  deleteVm(vmId: string): Observable<any> {
    return this.http.delete(`/v1/vms/${vmId}`);
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

  getInstallationList(vmId: number): Observable<Array<Installation>> {
    return this.http.get<Array<Installation>>(`/v1/installations`, {
      params: {
        id: vmId.toString()
      }
    }).pipe(map(res => plainToClass(Installation, res)));
  }

  createInstallation(vmId: number, packageName, packageTag: string): Observable<any> {
    return this.http.post(`/v1/installations/${vmId}`, {
      package_name: packageName,
      package_tag: packageTag
    });
  }

  deleteInstallation(vmId: number, installation: Installation): Observable<any> {
    return this.http.delete(`/v1/installations/${vmId}`, {
      params: {
        pkg_name: installation.name,
        pkg_tag: installation.tag
      }
    });
  }
}
