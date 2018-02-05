import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-stdxxx',
  styleUrls: ['./stdxxx.component.css'],
  templateUrl: './stdxxx.component.html',
})
export class StdxxxComponent {

  @Input()
  title: string;

  @Input()
  readonly = false;

  @Output()
  contentChange = new EventEmitter<string>();

  private _content: string;

  @Input()
  set content(content: string) {
    this._content = content;
    this.contentChange.emit(content);
  }

  get content() { return this._content; }

}
