import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'stdxxx',
  template: `
    <div class="flex">
      <h2>Standard input</h2>
      <textarea></textarea>
    </div>
  `,
  styles: [
    `
    .flex {
      display: flex;
      flex-direction: column;
      height: 100%;
    }
    textarea {
      flex: 1;
      font-family: 'Lucida Grande', sans-serif;
      resize: none;
    }
    `
  ]
})
export class StdxxxComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
