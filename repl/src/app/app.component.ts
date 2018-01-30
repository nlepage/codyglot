import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <div class="grid">
      <toolbar></toolbar>
      <app-source></app-source>
      <stdxxx class="stdin"></stdxxx>
      <stdxxx class="stdout"></stdxxx>
      <stdxxx class="stderr"></stdxxx>
    </div>
  `,
  styles: [
    `
    .grid {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      grid-template-rows: 60px 2fr 1fr;
      height: 100vh;
    }
    toolbar {
      grid-column: 1 / 3;
      grid-row: 1;
    }
    app-source {
      grid-column: 1;
      grid-row: 2;
    }
    .stdin {
      border-top: 1px solid #370037;
      grid-column: 1;
      grid-row: 3;
    }
    .stdout {
      border-left: 1px solid #370037;
      grid-column: 2;
      grid-row: 2;
    }
    .stderr {
      border-left: 1px solid #370037;
      border-top: 1px solid #370037;
      grid-column: 2;
      grid-row: 3;
    }
    `,
  ]
})
export class AppComponent {
  title = 'app';
}
