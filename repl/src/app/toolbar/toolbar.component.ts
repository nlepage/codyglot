import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { get } from 'lodash';
import { LanguageInfo, LanguagesService } from '../languages.service';

@Component({
  selector: 'app-toolbar',
  styleUrls: ['./toolbar.component.css'],
  templateUrl: './toolbar.component.html',
})
export class ToolbarComponent implements OnInit {

  languages: LanguageInfo[];

  @Output() selectLanguage = new EventEmitter<LanguageInfo>();

  @Output() run = new EventEmitter<void>();

  @Input() executing: boolean;

  @Input() compilationTime: string;

  @Input() runningTime: string;

  private _language: LanguageInfo;

  constructor(private languagesService: LanguagesService) {}

  ngOnInit() {
    this.languagesService.languages.subscribe((languages) => {
      this.languages = languages;
      this.language = languages[0].key;
    });
  }

  get language() { return get(this._language, 'key'); }
  set language(key: string) {
    this._language = this.languages.find((language) => language.key === key);
    this.selectLanguage.emit(this._language);
  }

  runClick = () => this.run.emit();
}
