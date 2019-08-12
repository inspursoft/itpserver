import { Component, OnInit } from '@angular/core';
import { CompatibilityService } from '../../compatibility.service';
import { Vm } from '../../compatibility.type';
import { NzMessageService, NzModalService } from 'ng-zorro-antd';
import { VmDetailComponent } from '../vm-detail/vm-detail.component';
import { SharedActionService } from '../../../shared/shared.action.service';

export enum ViewModel {list, card}

@Component({
  selector: 'app-vm-list',
  templateUrl: './vm-list.component.html',
  styleUrls: ['./vm-list.component.less']
})
export class VmListComponent implements OnInit {
  vmList: Array<Vm>;
  loading = false;
  pageIndex = 1;
  pageSize = 10;
  deleteVmWIP = false;
  viewModel: ViewModel = ViewModel.list;

  constructor(private service: CompatibilityService,
              private modalService: NzModalService,
              private messageService: NzMessageService,
              private sharedAcitonService: SharedActionService) {
    this.vmList = Array<Vm>();
  }

  ngOnInit() {
    this.retrieve();
  }

  showListModel() {
    this.viewModel = ViewModel.list;
  }

  showCardModel() {
    this.viewModel = ViewModel.card;
  }

  retrieve() {
    this.loading = true;
    this.service.getVmList().subscribe(
      (res: Array<Vm>) => this.vmList = res,
      () => this.loading = false,
      () => this.loading = false
    );
  }

  deleteVm(vmName: string) {
    if (!this.deleteVmWIP) {
      this.modalService.confirm({
        nzTitle: '删除',
        nzContent: '<b style="color: red;">确定要删除该测试环境?</b>',
        nzOkText: 'Yes',
        nzOkType: 'danger',
        nzOkLoading: this.deleteVmWIP,
        nzOnOk: () => this.deleteAction(vmName),
        nzCancelText: 'No'
      });
    }
  }

  deleteAction(vmName: string) {
    this.deleteVmWIP = true;
    this.service.deleteVm(vmName).subscribe(
      () => this.messageService.success('删除测试环境成功！'),
      () => {
        this.messageService.warning('删除测试环境失败！');
        this.deleteVmWIP = false;
      },
      () => {
        this.retrieve();
        this.deleteVmWIP = false;
      }
    );
  }

  showDetailInfo(vmName: string) {
    if (!this.deleteVmWIP) {
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
  }

  createVm() {
    this.sharedAcitonService.createVm().subscribe((vm: Vm) => {
      if (vm) {
        this.retrieve();
      }
    });
  }
}
