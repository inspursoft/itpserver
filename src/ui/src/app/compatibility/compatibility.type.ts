import { Expose, Type } from 'class-transformer';

export class VmSpec {
  cpus: number;
  extras: string;
  memory: string;
  storage: string;
  vid: string;

  get postBody(): object {
    return {
      cpus: this.cpus,
      memory: `${this.memory}`,
      storage: `${this.storage}G`
    };
  }
}

export class Vm {
  id: number;
  @Expose({name: 'vm_ip'}) ip: string;
  @Expose({name: 'vm_name'}) name: string;
  @Expose({name: 'vm_os'}) os: string;
  @Type(() => VmSpec)
  @Expose({name: 'vm_spec'})
  spec: VmSpec;

  constructor() {
    this.spec = new VmSpec();
  }

  postBody(): object {
    return {
      vm_ip: this.ip,
      vm_name: this.name,
      vm_os: this.os,
      vm_spec: this.spec.postBody
    };
  }
}

export class Package {
  @Expose({name: 'package_id'}) id: number;
  @Expose({name: 'package_name'}) name: string;
  @Expose({name: 'package_tag'}) tag: string;
  @Expose({name: 'vm_name'}) vmName: string;
}

export class Installation {
  @Expose({name: 'package_name'}) name: string;
  @Expose({name: 'package_tag'}) tag: string;
}


export class BaseOs {
  name: string;
  version: string;
}
