import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd';

@Component({
  selector: 'app-create-package',
  templateUrl: './create-package.component.html',
  styleUrls: ['./create-package.component.less']
})
export class CreatePackageComponent implements OnInit {
  packageFromGroup: FormGroup;

  constructor(private fb: FormBuilder,
              private msg: NzMessageService) {
    this.packageFromGroup = this.fb.group({
      name: ['', [Validators.required]],
      tag: ['', [Validators.required]]
    });
  }

  ngOnInit() {
  }

  restForm() {
    this.packageFromGroup.reset();
    for (const key of Object.keys(this.packageFromGroup.controls)) {
      this.packageFromGroup.controls[key].markAsPristine();
      this.packageFromGroup.controls[key].updateValueAndValidity();
    }
  }

  handleChange({ file, fileList }: { [key: string]: any }): void {
    const status = file.status;
    if (status !== 'uploading') {
      console.log(file, fileList);
    }
    if (status === 'done') {
      this.msg.success(`${file.name} file uploaded successfully.`);
    } else if (status === 'error') {
      this.msg.error(`${file.name} file upload failed.`);
    }
  }
}
