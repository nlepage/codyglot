import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

import { LanguageInfo } from '../languages.service';

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

  @Input()
  language: LanguageInfo;

  @Output()
  sourceChange = new EventEmitter<string>();

  private _source: string;

  set source(source: string) {
    this._source = source;
    this.sourceChange.emit(source);
  }

  get source() { return this._source; }

}
