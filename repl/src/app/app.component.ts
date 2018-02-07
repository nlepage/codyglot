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
  result: ExecuteResult;
  executing = false;

  constructor(private executeService: ExecuteService) {
    this.executeService.result.subscribe((result) => {
      this.result = result;
      this.executing = false;
    });
  }

  execute = ({ language }: { language: string }) => {
    this.executing = true;
    this.result = undefined;
    this.executeService.execute(language, this.source, this.stdin);
  }
}
