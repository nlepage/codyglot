import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'stdxxx',
  templateUrl: './stdxxx.component.html',
  styleUrls: ['./stdxxx.component.css']
})
export class StdxxxComponent implements OnInit {

  @Input()
  title: string

  @Input()
  readonly = false

  @Input()
  content: string

  constructor() { }

  ngOnInit() {
  }

}
