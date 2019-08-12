import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-log-viewer',
  templateUrl: './log-viewer.component.html',
  styleUrls: ['./log-viewer.component.less']
})
export class LogViewerComponent implements OnInit {
  @Input() logs: string;
  arrLogs: Array<string>;

  constructor() {
    this.arrLogs = Array<string>();
  }

  ngOnInit() {
    if (this.logs) {
      this.arrLogs = this.logs.split(/\n/);
    }
  }

}
