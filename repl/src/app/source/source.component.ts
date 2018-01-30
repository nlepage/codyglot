import { Component, OnInit } from '@angular/core';
import 'brace';
import 'brace/theme/chrome';
import 'brace/mode/golang';
import 'brace/mode/javascript';

@Component({
  selector: 'app-source',
  template: `
    <div class="flex">
      <h2>Source code</h2>
      <ace-editor mode="javascript" theme="chrome"></ace-editor>
    </div>
  `,
  styles: [
    `
    .flex {
      display: flex;
      flex-direction: column;
      height: 100%;
    }
    h2 {
      background-color: #a66fa6;
      color: #fff;
      font-size: 16px;
      margin: 0;
      padding: 5px;
    }
    ace-editor {
      flex: 1;
      font-size: 16px;
    }
    `
  ]
})
export class SourceComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
