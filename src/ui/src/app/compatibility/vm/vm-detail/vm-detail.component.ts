import {Component, Input, OnInit} from '@angular/core';
import {Installation, Vm} from '../../compatibility.type';
import {CompatibilityService} from '../../compatibility.service';
import {NzModalRef} from 'ng-zorro-antd';

@Component({
  selector: 'app-installation-detail',
  templateUrl: './vm-detail.component.html',
  styleUrls: ['./vm-detail.component.less']
})
export class VmDetailComponent implements OnInit {
  @Input() vmName: string;
  installationList: Array<Installation>;
  loading = false;
  pageIndex = 1;
  pageSize = 10;

  constructor(private service: CompatibilityService,
              private modal: NzModalRef) {
    this.installationList = Array<Installation>();
  }

  ngOnInit() {
    this.retrieve();
  }

  retrieve() {
    this.loading = true;
    this.service.getInstallationList(this.vmName).subscribe(
      (res: Array<Installation>) => this.installationList = res,
      () => this.loading = false,
      () => this.loading = false
    );
  }

  deleteInstallation(installation: Installation) {
    this.service.deleteInstallation(this.vmName, installation).subscribe(
      () => this.retrieve(),
      () => this.modal.destroy()
    );
  }
}
