import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';

@Component({
  selector: 'stdxxx',
  templateUrl: './stdxxx.component.html',
  styleUrls: ['./stdxxx.component.css']
})
export class StdxxxComponent implements OnInit {

  @Input()
  title: string;

  @Input()
  readonly = false;

  private _content: string;

  @Output() onContentChange = new EventEmitter<string>();

  constructor() {}

  ngOnInit() {
  }

  @Input()
  set content(content: string) {
    this._content = content;
    this.onContentChange.emit(content);
  }

  get content() { return this._content; }

}
