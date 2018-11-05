import { Component } from '@angular/core';

import { ExecuteService } from './execute.service';

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

  get stdout() {
    return this.getStdxxx('stdout')
  }

  get stderr() {
    return this.getStdxxx('stderr')
  }

  getStdxxx(prop: string) {
    if (!this.executeService.result) return undefined
    const { compilation, executions: [execution] } = this.executeService.result
    if (compilation && compilation.status) return compilation[prop]
    return execution[prop]
  }

  get result() { return this.executeService.result; }
}
