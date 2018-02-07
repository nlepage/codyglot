import { Component, EventEmitter, Input, Output } from '@angular/core';
import { get } from 'lodash';
import { LanguageInfo, LanguagesService } from '../languages.service';

@Component({
  selector: 'app-toolbar',
  styleUrls: ['./toolbar.component.css'],
  templateUrl: './toolbar.component.html',
})
export class ToolbarComponent {

  @Output() run = new EventEmitter<{language: string}>();

  @Input() executing: boolean;

  @Input() compilationTime: string;

  @Input() runningTime: string;

  constructor(private languagesService: LanguagesService) {}

  get language() { return this.languagesService.language; }
  set language(key: string) { this.languagesService.language = key; }

  get languages() { return this.languagesService.languages; }

  runClick = () => this.run.emit({ language: this.language });
}
