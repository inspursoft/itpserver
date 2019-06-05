import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ModalButtonOptions, NzMessageService, NzModalRef, NzModalService } from 'ng-zorro-antd';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { CreateVmComponent } from './create-vm/create-vm.component';
import { Vm } from '../compatibility/compatibility.type';
import { CreatePackageComponent } from './create-package/create-package.component';
import { LogViewerComponent } from './log-viewer/log-viewer.component';

@Injectable({providedIn: 'root'})
export class SharedActionService {
  private vmModalRef: NzModalRef;
  private packageModalRef: NzModalRef;
  private logViewerModalRef: NzModalRef;

  private createVmAction(instance: CreateVmComponent) {
    if (instance.vmFormGroup.valid) {
      const newVm = new Vm();
      const restButtonOption: ModalButtonOptions = this.vmModalRef.getInstance().nzFooter[0];
      const cancelButtonOption: ModalButtonOptions = this.vmModalRef.getInstance().nzFooter[1];
      const createButtonOption: ModalButtonOptions = this.vmModalRef.getInstance().nzFooter[2];
      newVm.os = instance.vmFormGroup.get('vmOs').value;
      newVm.ip = instance.vmFormGroup.get('vmIp').value;
      newVm.name = instance.vmFormGroup.get('vmName').value;
      newVm.spec.cpus = instance.vmFormGroup.get('vmCpu').value;
      newVm.spec.memory = instance.vmFormGroup.get('vmMemory').value;
      newVm.spec.storage = instance.vmFormGroup.get('vmStorage').value;
      restButtonOption.disabled = true;
      cancelButtonOption.disabled = true;
      createButtonOption.loading = true;
      instance.setDisabled(true);
      this.http.post(`/v1/vms`, newVm.postBody(), {responseType: 'text'}).subscribe(
        (res: string) => this.vmModalRef.close(res),
        (err: HttpErrorResponse) => {
          if (err.status === 409) {
            this.messageService.error('测试环境ID已经存在');
          } else {
            this.messageService.error('创建失败');
          }
          restButtonOption.disabled = false;
          cancelButtonOption.disabled = false;
          createButtonOption.loading = false;
          instance.setDisabled(false);
        }
      );
    } else {
      this.messageService.error('填写完整表单！');
    }
  }

  private createPackageAction(instance: CreatePackageComponent) {
    if (instance.packageFromGroup.valid && instance.fileList.length > 0) {
      const restButtonOption: ModalButtonOptions = this.packageModalRef.getInstance().nzFooter[0];
      const cancelButtonOption: ModalButtonOptions = this.packageModalRef.getInstance().nzFooter[1];
      const createButtonOption: ModalButtonOptions = this.packageModalRef.getInstance().nzFooter[2];
      const formData = new FormData();
      let vmName = instance.packageFromGroup.get('vmName').value;
      vmName = instance.packageFromGroup.get('vmName').value;
      formData.append('pkg', instance.fileList[0] as any);
      restButtonOption.disabled = true;
      cancelButtonOption.disabled = true;
      createButtonOption.loading = true;
      instance.setDisabled(true);
      this.http.post(`/v1/packages`, formData, {
        responseType: 'text',
        params: {vm_name: vmName}
      }).subscribe(
        (res: string) => this.packageModalRef.close(res),
        (err: HttpErrorResponse) => {
          if (err.status === 409) {
            this.messageService.error('Package已经存在');
          } else {
            this.messageService.error('创建失败');
          }
          restButtonOption.disabled = false;
          cancelButtonOption.disabled = false;
          createButtonOption.loading = false;
          instance.setDisabled(false);
        }
      );
    } else {
      this.messageService.error('填写完整表单！');
    }
  }

  constructor(private modalService: NzModalService,
              private messageService: NzMessageService,
              private http: HttpClient) {
  }

  createVm(): Observable<any> {
    this.vmModalRef = this.modalService.create({
      nzTitle: `创建测试环境`,
      nzContent: CreateVmComponent,
      nzClosable: false,
      nzMaskClosable: false,
      nzFooter: [{
        label: '重置',
        shape: 'default',
        onClick: (instance: CreateVmComponent) => instance.restForm()
      }, {
        label: '取消',
        shape: 'default',
        onClick: () => this.vmModalRef.destroy()
      }, {
        label: '创建',
        shape: 'primary',
        onClick: (instance: CreateVmComponent) => this.createVmAction(instance)
      }]
    });
    return this.vmModalRef.afterClose;
  }

  createPackage(): Observable<any> {
    this.packageModalRef = this.modalService.create({
      nzTitle: `创建工具集`,
      nzContent: CreatePackageComponent,
      nzClosable: false,
      nzMaskClosable: false,
      nzFooter: [{
        label: '重置',
        shape: 'default',
        onClick: (instance: CreatePackageComponent) => instance.restForm()
      }, {
        label: '取消',
        shape: 'default',
        onClick: () => this.packageModalRef.destroy()
      }, {
        label: '创建',
        shape: 'primary',
        onClick: (instance: CreatePackageComponent) => this.createPackageAction(instance)
      }]
    });
    return this.packageModalRef.afterClose;
  }

  createLogsViewer(logs: string): Observable<any> {
    this.logViewerModalRef = this.modalService.create({
      nzTitle: `Logs`,
      nzContent: LogViewerComponent,
      nzClosable: false,
      nzMaskClosable: false,
      nzWidth: 960,
      nzComponentParams: {logs},
      nzFooter: [{
        label: '关闭',
        shape: 'primary',
        onClick: () => this.logViewerModalRef.destroy()
      }]
    });
    return this.logViewerModalRef.afterClose;
  }
}
