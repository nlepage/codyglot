import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'toolbar',
  template: `
    <div class="flex">
      <h1>Codyglot REPL</h1>
      <div class="language pure-form">
        <select>
          <option>Go(Lang)</option>
          <option>NodeJS</option>
        </select>
      </div>
      <button class="pure-button run">Run</button>
    </div>
  `,
  styles: [
    `
    .flex {
      align-items: center;
      background-color: #530e53;
      display: flex;
      flex-direction: row;
      height: 100%;
    }
    h1 {
      color: #fff;
      font-size: 20px;
      margin-left: 20px;
    }
    .language {
      margin-left: 20px;
    }
    .run {
      background-color: #8a458a;
      color: #fff;
      margin-left: 20px;
    }
    `
  ]
})
export class ToolbarComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
