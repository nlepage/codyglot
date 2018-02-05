import { Component, EventEmitter, Input, OnInit, Output } from "@angular/core";

import { ILanguageInfo } from "../languages.service";

import "brace";
import "brace/mode/golang";
import "brace/mode/javascript";
import "brace/theme/chrome";

@Component({
  selector: "app-source",
  styleUrls: ["./source.component.css"],
  templateUrl: "./source.component.html",
})
export class SourceComponent {

  @Input()
  public language: ILanguageInfo;

  @Output()
  public onSourceChange = new EventEmitter<string>();

  private _source: string;

  set source(source: string) {
    this._source = source;
    this.onSourceChange.emit(source);
  }

  get source() { return this._source; }

}
