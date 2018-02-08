import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

import { LanguagesService } from '../languages.service';

import 'brace';
import 'brace/mode/golang';
import 'brace/mode/javascript';
import 'brace/mode/typescript';
import 'brace/theme/chrome';

@Component({
  selector: 'app-source',
  styleUrls: ['./source.component.css'],
  templateUrl: './source.component.html',
})
export class SourceComponent {

  @Output()
  sourceChange = new EventEmitter<string>();

  private _source: string;

  constructor(private languagesService: LanguagesService) {}

  get source() { return this._source; }
  set source(source: string) {
    this._source = source;
    this.sourceChange.emit(source);
  }

  get language() { return this.languagesService.languageInfo; }
}
