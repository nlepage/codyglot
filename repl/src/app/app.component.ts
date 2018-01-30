import { Component } from '@angular/core';
import { LanguageInfo } from './languages.service'
import { ExecuteService } from './execute.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  language: LanguageInfo;
  source: string;
  stdin: string;
  stdout: string;
  stderr: string;

  constructor(private executeService: ExecuteService) {
    this.executeService.result.subscribe(({ stdout, stderr }) => {
      this.stdout = stdout;
      this.stderr = stderr;
    })
  }

  execute = () => {
    this.executeService.execute(this.language.key, this.source, this.stdin);
  }
}
