import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { InstallationListComponent } from './installation-list.component';
import { SharedModule } from '../../../shared/shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterModule } from '@angular/router';
import { RouteInstallation } from '../../../shared/shared.const';

describe('InstallationListComponent', () => {
  let component: InstallationListComponent;
  let fixture: ComponentFixture<InstallationListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [InstallationListComponent],
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule,
        RouterModule.forChild([
          {path: RouteInstallation, component: InstallationListComponent}
        ])],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InstallationListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
