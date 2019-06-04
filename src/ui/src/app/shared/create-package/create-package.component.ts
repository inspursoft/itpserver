import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, ValidationErrors, Validators } from '@angular/forms';
import { NzMessageService, UploadFile, UploadFilter } from 'ng-zorro-antd';
import { Observable, of } from 'rxjs';
import { Vm } from '../../compatibility/compatibility.type';
import { SharedService } from '../shared.service';

@Component({
  selector: 'app-create-package',
  templateUrl: './create-package.component.html',
  styleUrls: ['./create-package.component.less']
})
export class CreatePackageComponent implements OnInit {
  packageFromGroup: FormGroup;
  fileList: Array<UploadFile>;
  uploadFilter: Array<UploadFilter>;
  createPackageWip = false;
  vmList: Array<Vm>;

  constructor(private fb: FormBuilder,
              private sharedService: SharedService,
              private msg: NzMessageService) {
    this.fileList = Array<UploadFile>();
    this.uploadFilter = Array<UploadFilter>();
    this.vmList = Array<Vm>();
    this.packageFromGroup = this.fb.group({
      // name: ['', [Validators.required], [this.packageNameValidator]],
      // tag: ['', [Validators.required]],
      vmName: ['', [Validators.required]]
    });
  }

  ngOnInit() {
    this.uploadFilter = [{
      name: 'type', fn: (fileList: Array<UploadFile>) => {
        const filterFiles = fileList.filter(w => ['application/zip'].indexOf(w.type) >= 0);
        if (filterFiles.length !== fileList.length) {
          this.msg.error(`包含文件格式不正确，只支持 'zip' 格式`);
          return filterFiles;
        }
        return fileList;
      }
    }];
    this.sharedService.getVmList().subscribe((vmList: Array<Vm>) => this.vmList = vmList);
  }

  setDisabled(disabled: boolean) {
    this.createPackageWip = disabled;
    if (disabled) {
      for (const key of Object.keys(this.packageFromGroup.controls)) {
        (Reflect.get(this.packageFromGroup.controls, key) as FormControl).disable({onlySelf: true});
      }
    } else {
      for (const key of Object.keys(this.packageFromGroup.controls)) {
        (Reflect.get(this.packageFromGroup.controls, key) as FormControl).enable({onlySelf: true});
      }
    }
  }

  packageNameValidator(control: AbstractControl): Promise<ValidationErrors | null> | Observable<ValidationErrors | null> {
    /*TODO:check package name whether exists*/
    return of(null);
  }

  restForm() {
    this.packageFromGroup.reset();
    for (const key of Object.keys(this.packageFromGroup.controls)) {
      this.packageFromGroup.controls[key].markAsPristine();
      this.packageFromGroup.controls[key].updateValueAndValidity();
    }
  }

  beforeUpload = (file: UploadFile) => {
    this.fileList.splice(0, this.fileList.length);
    this.fileList = this.fileList.concat(file);
    return false;
  };
}
