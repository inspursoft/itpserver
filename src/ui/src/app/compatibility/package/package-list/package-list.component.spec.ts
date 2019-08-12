import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { PackageListComponent } from './package-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { RouterModule } from '@angular/router';
import { RoutePackage } from '../../../shared/shared.const';

describe('PackageListComponent', () => {
  let component: PackageListComponent;
  let fixture: ComponentFixture<PackageListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [PackageListComponent],
      imports: [
        SharedModule,
        RouterModule.forChild([
          {path: RoutePackage, component: PackageListComponent}
        ])
      ]
    }).compileComponents();
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
