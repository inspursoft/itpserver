import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VmListComponent } from './vm-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterModule } from '@angular/router';
import { RouteVm } from '../../../shared/shared.const';

describe('VmListComponent', () => {
  let component: VmListComponent;
  let fixture: ComponentFixture<VmListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [VmListComponent],
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule,
        RouterModule.forChild([
          {path: RouteVm, component: VmListComponent}
        ])]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VmListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
