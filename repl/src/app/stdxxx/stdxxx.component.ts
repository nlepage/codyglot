import { Component, EventEmitter, Input, Output } from "@angular/core";

@Component({
  selector: "stdxxx",
  styleUrls: ["./stdxxx.component.css"],
  templateUrl: "./stdxxx.component.html",
})
export class StdxxxComponent {

  @Input()
  public title: string;

  @Input()
  public readonly = false;

  @Output()
  public onContentChange = new EventEmitter<string>();

  private _content: string;

  @Input()
  set content(content: string) {
    this._content = content;
    this.onContentChange.emit(content);
  }

  get content() { return this._content; }

}
