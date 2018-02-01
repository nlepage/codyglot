import { Component, OnInit, OnDestroy } from '@angular/core';

@Component({
  selector: 'loading',
  templateUrl: './loading.component.html'
})
export class LoadingComponent implements OnInit, OnDestroy {

  interval: number

  angle = 0

  constructor() {}

  ngOnInit() {
    this.interval = window.setInterval(() => {
      this.angle = (this.angle + 0.1) % 1;
      console.log(this.angle)
    }, 50)
  }

  ngOnDestroy() {
    window.clearInterval(this.interval)
  }

  get rotate() {
    return `rotate(${this.angle}turn)`
  }
}
