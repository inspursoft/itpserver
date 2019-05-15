import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { BaseOsListComponent } from './base-os-list.component';
import { SharedModule } from '../../shared/shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RouterModule } from '@angular/router';
import { RouteBaseOs,} from '../../shared/shared.const';

describe('BaseOsListComponent', () => {
  let component: BaseOsListComponent;
  let fixture: ComponentFixture<BaseOsListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [BaseOsListComponent],
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule,
        RouterModule.forChild([
          {path: RouteBaseOs, component: BaseOsListComponent},
        ])]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BaseOsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
