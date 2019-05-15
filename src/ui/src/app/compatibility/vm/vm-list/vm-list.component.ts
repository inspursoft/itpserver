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

  showListModel(){
    this.viewModel = ViewModel.list;
  }

  showCardModel(){
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

  deleteVm(vmId: string) {
    this.modalService.confirm({
      nzTitle: '删除',
      nzContent: '<b style="color: red;">确定要删除该测试环境?</b>',
      nzOkText: 'Yes',
      nzOkType: 'danger',
      nzOnOk: () => this.deleteAction(vmId),
      nzCancelText: 'No'
    });
  }

  deleteAction(vmId: string) {
    this.service.deleteVm(vmId).subscribe(
      () => this.messageService.success('删除测试环境成功！'),
      () => this.messageService.warning('删除测试环境失败！'),
      () => this.retrieve()
    );
  }

  showDetailInfo(vmId: number) {
    const modal = this.modalService.create({
      nzTitle: '详细信息',
      nzContent: VmDetailComponent,
      nzComponentParams: {Id: vmId},
      nzFooter: [{
        label: '确定',
        shape: 'primary',
        onClick: () => modal.destroy()
      }]
    });
  }

  createVm() {
    this.sharedAcitonService.createVm().subscribe((vm: Vm) => {
      if (vm) {
        this.retrieve();
      }
    });
  }
}
