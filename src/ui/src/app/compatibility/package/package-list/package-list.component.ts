import { Component, OnInit } from '@angular/core';
import { Package } from '../../compatibility.type';
import { CompatibilityService } from '../../compatibility.service';
import { NzMessageService, NzModalService } from 'ng-zorro-antd';
import { SharedActionService } from '../../../shared/shared.action.service';

@Component({
  selector: 'app-package-list',
  templateUrl: './package-list.component.html',
  styleUrls: ['./package-list.component.less']
})
export class PackageListComponent implements OnInit {
  packageList: Array<Package>;
  pageIndex = 1;
  pageSize = 10;
  loading = false;

  constructor(private service: CompatibilityService,
              private messageService: NzMessageService,
              private modalService: NzModalService,
              private sharedActionService: SharedActionService) {
    this.packageList = Array<Package>();
  }

  ngOnInit() {
    this.retrieve();
  }

  retrieve() {
    this.loading = true;
    this.service.getPackageList().subscribe(
      (res: Array<Package>) => this.packageList = res,
      () => this.loading = false,
      () => this.loading = false
    );
  }

  deletePackage(packageName, packageTag: string) {
    this.modalService.confirm({
      nzTitle: '删除',
      nzContent: '<b style="color: red;">确定要删除该工具集?</b>',
      nzOkText: 'Yes',
      nzOkType: 'danger',
      nzOnOk: () => this.deleteAction(packageName, packageTag),
      nzCancelText: 'No'
    });
  }

  deleteAction(packageName, packageTag: string) {
    this.service.deletePackage(packageName, packageTag).subscribe(
      () => this.messageService.success('删除工具集成功！'),
      () => this.messageService.warning('删除工具集失败！'),
      () => this.retrieve()
    );
  }

  createPackage() {
    this.sharedActionService.createPackage().subscribe((logs: string) => {
      if (logs) {
        this.sharedActionService.createLogsViewer(logs).subscribe(() => this.retrieve());
      }
    });
  }
}
