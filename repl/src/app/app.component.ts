import { Component } from '@angular/core';

import { ExecuteService, ExecuteResult } from './execute.service';

@Component({
  selector: 'app-root',
  styleUrls: ['./app.component.css'],
  templateUrl: './app.component.html',
})
export class AppComponent {

  source: string;
  stdin: string;

  constructor(private executeService: ExecuteService) {}

  execute = ({ language }: { language: string }) => {
    this.executeService.execute(language, this.source, this.stdin);
  }

  get result() { return this.executeService.result; }
}
