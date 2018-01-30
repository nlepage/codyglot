import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ReplaySubject } from 'rxjs';
import 'rxjs/add/operator/map';
import { keyBy, sortBy } from 'lodash';

export type LanguageInfo = {
  key: string,
  name: string,
  mode: string,
}

@Injectable()
export class LanguageService {

  private languagesInfo = keyBy([
    { key: 'nodejs', name: 'NodeJS', mode: 'javascript' },
    { key: 'golang', name: 'Go(lang)', mode: 'golang' },
  ], 'key')

  private _languages = new ReplaySubject<LanguageInfo[]>()

  constructor(private http: HttpClient) {
    this.http.get<{languages:String[]}>('/api/languages')
      .map(res => res.languages)
      .map(languages => languages.map(this.getLanguageInfo))
      .map(languages => sortBy(languages, 'name'))
      .subscribe(this._languages);
  }

  getLanguageInfo = (key: string) => {
    if (this.languagesInfo[key]) return this.languagesInfo[key];
    return {
      key,
      name: key,
      mode: key,
    };
  }

  get languages() { return this._languages.asObservable(); }

}
