import { Component } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, ValidationErrors, Validators } from '@angular/forms';
import { Observable, of } from 'rxjs';

@Component({
  selector: 'app-create-vm',
  templateUrl: './create-vm.component.html',
  styleUrls: ['./create-vm.component.less']
})
export class CreateVmComponent {
  vmFormGroup: FormGroup;

  constructor(private fb: FormBuilder) {
    this.vmFormGroup = this.fb.group({
      vmIp: ['', [Validators.required], [this.vmIpValidator]],
      vmName: ['', [Validators.required]],
      vmOs: ['', [Validators.required]],
      vmCpu: [1, [Validators.required, Validators.max(16), Validators.min(1)]],
      vmMemory: [1024, [Validators.required, Validators.min(1024)]],
      vmStorage: [1, [Validators.required, Validators.min(1)]]
    });
  }

  setDisabled(disabled: boolean) {
    if (disabled) {
      for (const key of Object.keys(this.vmFormGroup.controls)) {
        (Reflect.get(this.vmFormGroup.controls, key) as FormControl).disable({onlySelf: true});
      }
    } else {
      for (const key of Object.keys(this.vmFormGroup.controls)) {
        (Reflect.get(this.vmFormGroup.controls, key) as FormControl).enable({onlySelf: true});
      }
    }
  }

  vmIpValidator(control: AbstractControl): Promise<ValidationErrors | null> | Observable<ValidationErrors | null> {
    /*TODO: check ip address whether exists.*/
    return of(null);
  }

  restForm() {
    this.vmFormGroup.reset({vmCpu: 1, vmMemory: 1024, vmStorage: 1});
    for (const key of Object.keys(this.vmFormGroup.controls)) {
      this.vmFormGroup.controls[key].markAsPristine();
      this.vmFormGroup.controls[key].updateValueAndValidity();
    }
  }
}
