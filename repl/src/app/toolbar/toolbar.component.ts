import { Component, EventEmitter, Output } from '@angular/core';
import { get } from 'lodash';
import { ExecuteService } from '../execute.service';
import { LanguageInfo, LanguagesService } from '../languages.service';

@Component({
  selector: 'app-toolbar',
  styleUrls: ['./toolbar.component.css'],
  templateUrl: './toolbar.component.html',
})
export class ToolbarComponent {

  @Output() run = new EventEmitter<{language: string}>();

  constructor(private languagesService: LanguagesService, private executeService: ExecuteService) {}

  runClick = () => this.run.emit({ language: this.language });

  get language() { return this.languagesService.language; }
  set language(key: string) { this.languagesService.language = key; }

  get languages() { return this.languagesService.languages; }

  get result() { return this.executeService.result; }

  get execution() { return this.result && this.result.executions && this.result.executions[0] }

  get executing() { return this.executeService.executing; }
}
