import { Component } from '@angular/core';
import { LanguageInfo } from './languages.service'
import { ExecuteResult, ExecuteService } from './execute.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  language: LanguageInfo;
  source: string;
  stdin: string;
  result: ExecuteResult;
  executing = false;

  constructor(private executeService: ExecuteService) {
    this.executeService.result.subscribe(result => {
      this.result = result;
      this.executing = false;
    });
  }

  execute = () => {
    this.executing = true
    this.result = undefined
    this.executeService.execute(this.language.key, this.source, this.stdin);
  }
}
