import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Vm } from '../compatibility/compatibility.type';
import { Observable } from 'rxjs';
import { map, tap } from 'rxjs/operators';
import { plainToClass } from 'class-transformer';

@Injectable({
  providedIn: 'root'
})
export class SharedService {

  constructor(private http: HttpClient) {
  }

  createVm(vm: Vm): Observable<any> {
    return this.http.post(`/v1/vms`, vm.postBody());
  }

  getVmList(): Observable<any> {
    return this.http.get(`/v1/vms`);
  }
}
