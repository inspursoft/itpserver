import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { PackageListComponent } from './package-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { RouterModule } from '@angular/router';
import { RoutePackage } from '../../../shared/shared.const';
import { NZ_ICONS } from 'ng-zorro-antd';
import { AccountBookFill, AlertFill, AlertOutline } from '@ant-design/icons-angular/icons';
import { IconDefinition } from '@ant-design/icons-angular';

const icons: IconDefinition[] = [AccountBookFill, AlertOutline, AlertFill];

describe('PackageListComponent', () => {
  let component: PackageListComponent;
  let fixture: ComponentFixture<PackageListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [PackageListComponent],
      providers: [{provide: NZ_ICONS, useValue: icons}],
      imports: [
        SharedModule,
        RouterModule.forChild([
          {path: RoutePackage, component: PackageListComponent}
        ])
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PackageListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
