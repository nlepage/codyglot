import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { LanguageInfo } from '../languages.service'
import 'brace';
import 'brace/theme/chrome';
import 'brace/mode/golang';
import 'brace/mode/javascript';

@Component({
  selector: 'app-source',
  templateUrl: './source.component.html',
  styleUrls: ['./source.component.css']
})
export class SourceComponent implements OnInit {

  @Input()
  language: LanguageInfo;

  private _source: string;

  @Output()
  onSourceChange = new EventEmitter<string>();

  constructor() {}

  ngOnInit() {
  }

  set source(source: string) {
    this._source = source;
    this.onSourceChange.emit(source);
  }

  get source() { return this._source; }

}
