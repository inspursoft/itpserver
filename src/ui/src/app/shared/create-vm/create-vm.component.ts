import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-create-vm',
  templateUrl: './create-vm.component.html',
  styleUrls: ['./create-vm.component.less']
})
export class CreateVmComponent {
  vmFormGroup: FormGroup;

  constructor(private fb: FormBuilder) {
    this.vmFormGroup = this.fb.group({
      vmIp: ['', [Validators.required]],
      vmName: ['', [Validators.required]],
      vmOs: ['', [Validators.required]],
      vmCpu: [1, [Validators.required, Validators.max(16), Validators.min(1)]],
      vmMemory: [1024, [Validators.required, Validators.min(1024)]],
      vmStorage: [1, [Validators.required, Validators.min(1)]]
    });
  }

  restForm() {
    this.vmFormGroup.reset({vmCpu: 1, vmMemory: 1024, vmStorage: 1});
    for (const key of Object.keys(this.vmFormGroup.controls)) {
      this.vmFormGroup.controls[key].markAsPristine();
      this.vmFormGroup.controls[key].updateValueAndValidity();
    }
  }
}
