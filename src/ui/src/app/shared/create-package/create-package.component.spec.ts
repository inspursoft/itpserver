import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { CreatePackageComponent } from './create-package.component';
import { SharedModule } from '../shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('CreatePackageComponent', () => {
  let component: CreatePackageComponent;
  let fixture: ComponentFixture<CreatePackageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreatePackageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
