import { Component, EventEmitter, Input, OnInit, Output } from "@angular/core";
import { get } from "lodash";
import { ILanguageInfo, LanguagesService } from "../languages.service";

@Component({
  selector: "toolbar",
  styleUrls: ["./toolbar.component.css"],
  templateUrl: "./toolbar.component.html",
})
export class ToolbarComponent implements OnInit {

  public languages: ILanguageInfo[];

  @Output() public onSelectLanguage = new EventEmitter<ILanguageInfo>();

  @Output() public onRun = new EventEmitter<void>();

  @Input() public executing: boolean;

  @Input() public compilationTime: string;

  @Input() public runningTime: string;

  private _language: ILanguageInfo;

  constructor(private languagesService: LanguagesService) {}

  public ngOnInit() {
    this.languagesService.languages.subscribe((languages) => {
      this.languages = languages;
      this.language = languages[0].key;
    });
  }

  get language() { return get(this._language, "key"); }
  set language(key: string) {
    this._language = this.languages.find((language) => language.key === key);
    this.onSelectLanguage.emit(this._language);
  }

  public run = () => this.onRun.emit();
}
