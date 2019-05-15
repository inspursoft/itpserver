import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import { RouteCompatibility, RouteInstallation, RouteVm } from '../shared/shared.const';

@Component({
  selector: 'app-main-content',
  templateUrl: './main-content.component.html',
  styleUrls: ['./main-content.component.less']
})
export class MainContentComponent implements OnInit {

  constructor(private router: Router) {
  }

  ngOnInit() {
    this.router.navigate([`/${RouteCompatibility}/${RouteInstallation}`]).then();
  }

}
