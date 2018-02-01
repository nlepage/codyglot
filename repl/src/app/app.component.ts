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

  constructor(private executeService: ExecuteService) {
    this.executeService.result.subscribe(result => this.result = result);
  }

  execute = () => {
    this.executeService.execute(this.language.key, this.source, this.stdin);
  }
}
