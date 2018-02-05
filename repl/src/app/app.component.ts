import { Component } from "@angular/core";

import { ExecuteService, IExecuteResult } from "./execute.service";
import { ILanguageInfo } from "./languages.service";

@Component({
  selector: "app-root",
  styleUrls: ["./app.component.css"],
  templateUrl: "./app.component.html",
})
export class AppComponent {

  public language: ILanguageInfo;
  public source: string;
  public stdin: string;
  public result: IExecuteResult;
  public executing = false;

  constructor(private executeService: ExecuteService) {
    this.executeService.result.subscribe((result) => {
      this.result = result;
      this.executing = false;
    });
  }

  public execute = () => {
    this.executing = true;
    this.result = undefined;
    this.executeService.execute(this.language.key, this.source, this.stdin);
  }
}
