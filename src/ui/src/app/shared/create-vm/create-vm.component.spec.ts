import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { CreateVmComponent } from './create-vm.component';
import { SharedModule } from '../shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('CreateVmComponent', () => {
  let component: CreateVmComponent;
  let fixture: ComponentFixture<CreateVmComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateVmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
