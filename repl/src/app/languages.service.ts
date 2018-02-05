import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";

import { ReplaySubject } from "rxjs";
import "rxjs/add/operator/map";

import { keyBy, sortBy } from "lodash";

export interface ILanguageInfo {
  key: string;
  name: string;
  mode: string;
}

@Injectable()
export class LanguagesService {

  private languagesInfo = keyBy([
    { key: "nodejs", name: "NodeJS", mode: "javascript" },
    { key: "golang", name: "Go(lang)", mode: "golang" },
  ], "key");

  private languages$ = new ReplaySubject<ILanguageInfo[]>();

  constructor(private http: HttpClient) {
    this.http.get<{languages: string[]}>("/api/languages")
      .map((res) => res.languages)
      .map((languages) => languages.map(this.getLanguageInfo))
      .map((languages) => sortBy(languages, "name"))
      .subscribe(this.languages$);
  }

  public getLanguageInfo = (key: string) => {
    if (this.languagesInfo[key]) {
      return this.languagesInfo[key];
    }
    return {
      key,
      mode: key,
      name: key,
    };
  }

  get languages() { return this.languages$.asObservable(); }
}
