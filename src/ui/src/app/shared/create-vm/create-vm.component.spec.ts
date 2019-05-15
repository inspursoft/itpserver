import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { CreateVmComponent } from './create-vm.component';
import { SharedModule } from '../shared.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NZ_ICONS } from 'ng-zorro-antd';
import { AccountBookFill, AlertFill, AlertOutline } from '@ant-design/icons-angular/icons';
import { IconDefinition } from '@ant-design/icons-angular';

const icons: IconDefinition[] = [AccountBookFill, AlertOutline, AlertFill];
describe('CreateVmComponent', () => {
  let component: CreateVmComponent;
  let fixture: ComponentFixture<CreateVmComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        BrowserModule,
        BrowserAnimationsModule,
        SharedModule],
      providers: [{provide: NZ_ICONS, useValue: icons}],
    })
      .compileComponents();
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
