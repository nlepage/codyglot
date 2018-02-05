import { Component } from '@angular/core';

import { ExecuteService, ExecuteResult } from './execute.service';
import { LanguageInfo } from './languages.service';

@Component({
  selector: 'app-root',
  styleUrls: ['./app.component.css'],
  templateUrl: './app.component.html',
})
export class AppComponent {

  language: LanguageInfo;
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

  execute = () => {
    this.executing = true;
    this.result = undefined;
    this.executeService.execute(this.language.key, this.source, this.stdin);
  }
}
