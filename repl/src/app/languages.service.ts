import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { ReplaySubject } from 'rxjs/ReplaySubject';
import 'rxjs/add/operator/map';

import { keyBy, sortBy } from 'lodash';

export interface LanguageInfo {
  key: string;
  name: string;
  mode: string;
}

@Injectable()
export class LanguagesService {

  private languagesInfo = keyBy([
    { key: 'golang', name: 'Go(lang)', mode: 'golang' },
    { key: 'javascript', name: 'JavaScript', mode: 'javascript' },
    { key: 'typescript', name: 'TypeScript', mode: 'typescript' },
  ], 'key');

  languages = new Array<LanguageInfo>();
  language: string;

  constructor(private http: HttpClient) {
    this.http.get<{languages: string[]}>('/api/languages')
      .map((res) => res.languages || new Array<string>('golang'))
      .map((languages) => languages.map(this.getLanguageInfo))
      .map((languages) => sortBy(languages, 'name'))
      .subscribe((languages) => {
        this.languages = languages;
        if (languages.length) {
          this.language = languages[0].key;
        }
      });
  }

  get languageInfo() { return this.getLanguageInfo(this.language); }

  private getLanguageInfo = (key: string): LanguageInfo => {
    if (this.languagesInfo[key]) {
      return this.languagesInfo[key];
    }
    return {
      key,
      mode: key,
      name: key,
    };
  }
}
