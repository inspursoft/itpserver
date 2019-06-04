import { Component, OnInit } from '@angular/core';
import { Package, Vm } from '../../compatibility.type';
import { CompatibilityService } from '../../compatibility.service';
import { SharedActionService } from '../../../shared/shared.action.service';
import { NzMessageService, NzModalService } from 'ng-zorro-antd';
import { VmDetailComponent } from '../../vm/vm-detail/vm-detail.component';

@Component({
  selector: 'app-installation-list',
  templateUrl: './installation-list.component.html',
  styleUrls: ['./installation-list.component.less']
})
export class InstallationListComponent implements OnInit {
  current = 0;
  vmList: Array<Vm>;
  packageList: Array<Package>;
  selectedVm: Vm;
  selectedPackage: Package;
  installationWip = false;

  constructor(private service: CompatibilityService,
              private messageService: NzMessageService,
              private sharedActionService: SharedActionService,
              private modalService: NzModalService) {
    this.vmList = Array<Vm>();
    this.packageList = Array<Package>();
  }

  ngOnInit(): void {
    this.service.getVmList().subscribe((res: Array<Vm>) => this.vmList = res);
    this.service.getPackageList().subscribe((res: Array<Package>) => this.packageList = res);
  }

  pre(): void {
    if (!this.installationWip) {
      this.current -= 1;
    }
  }

  next(): void {
    this.current += 1;
  }

  createInstallation() {
    if (!this.installationWip) {
      if (!this.selectedVm) {
        this.messageService.warning('请选择测试环境');
      } else if (!this.selectedPackage) {
        this.messageService.warning('请选择工具集');
      } else {
        this.installationWip = true;
        this.service.createInstallation(this.selectedVm.name, this.selectedPackage.name, this.selectedPackage.tag).subscribe(
          (logs: string) => {
            this.sharedActionService.createLogsViewer(logs);
            this.selectedPackage = null;
            this.selectedVm = null;
            this.current = 0;
            this.installationWip = false;
          },
          () => {
            this.messageService.warning('安装失败');
            this.installationWip = false;
          });
      }
    }
  }

  showDetailInfo(vmName: string, event: Event) {
    event.stopPropagation();
    const modal = this.modalService.create({
      nzTitle: '详细信息',
      nzContent: VmDetailComponent,
      nzComponentParams: {vmName},
      nzFooter: [{
        label: '确定',
        shape: 'primary',
        onClick: () => modal.destroy()
      }]
    });
  }

  selectVm(vm: Vm) {
    this.selectedVm = vm;
    this.current += 1;
  }

  selectPackage(packageInfo: Package) {
    this.selectedPackage = packageInfo;
    this.current += 1;
  }

  createPackage() {
    this.sharedActionService.createPackage().subscribe((logs: string) => {
      if (logs) {
        this.sharedActionService.createLogsViewer(logs).subscribe(() => {
          this.service.getPackageList().subscribe((res: Array<Package>) => this.packageList = res);
        });
      }
    });
  }

  createVm() {
    this.sharedActionService.createVm().subscribe((logs: string) => {
      if (logs) {
        this.sharedActionService.createLogsViewer(logs).subscribe(() => {
          this.service.getVmList().subscribe((res: Array<Vm>) => this.vmList = res);
        });
      }
    });
  }
}
