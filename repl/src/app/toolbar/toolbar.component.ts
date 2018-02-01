import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { LanguageInfo, LanguagesService } from '../languages.service';
import { get } from 'lodash';

@Component({
  selector: 'toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.css']
})
export class ToolbarComponent implements OnInit {

  languages: LanguageInfo[];

  private _language: LanguageInfo;

  @Output() onSelectLanguage = new EventEmitter<LanguageInfo>();

  @Output() onRun = new EventEmitter<void>();

  @Input() compilationTime: string;

  @Input() runningTime: string;

  constructor(private languagesService: LanguagesService) {}

  ngOnInit() {
    this.languagesService.languages.subscribe(languages => {
      this.languages = languages;
      this.language = languages[0].key;
    });
  }

  get language() { return get(this._language, 'key'); }
  set language(key: string) {
    this._language = this.languages.find(language => language.key === key);
    this.onSelectLanguage.emit(this._language);
  }

  run = () => this.onRun.emit();
}
