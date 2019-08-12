import { Component, OnInit } from '@angular/core';
import { BaseOs } from '../compatibility.type';
import { CompatibilityService } from '../compatibility.service';
import { NzMessageService, NzModalService } from 'ng-zorro-antd';

@Component({
  selector: 'app-base-os-list',
  templateUrl: './base-os-list.component.html',
  styleUrls: ['./base-os-list.component.less']
})
export class BaseOsListComponent implements OnInit {
  baseOsList: Array<BaseOs>;

  constructor(private service: CompatibilityService,
              private modalService: NzModalService,
              private messageService: NzMessageService) {
    this.baseOsList = Array<BaseOs>();
  }

  ngOnInit() {
    this.baseOsList.push(...[
      {name: 'Ubuntu', version: '7.0'},
      {name: 'CentOs', version: '1.0'}
    ]);
  }

  deleteBaseOs(os: BaseOs) {
    this.modalService.confirm({
      nzTitle: '删除',
      nzContent: '<b style="color: red;">确定要删除该基础镜像?</b>',
      nzOkText: 'Yes',
      nzOkType: 'danger',
      nzOnOk: () => this.messageService.warning('删除基础镜像失败！'),
      nzCancelText: 'No'
    });
  }
}
