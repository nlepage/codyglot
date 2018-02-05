import { Component, OnDestroy, OnInit } from '@angular/core';

@Component({
  selector: 'app-loading',
  templateUrl: './loading.component.html',
})
export class LoadingComponent implements OnInit, OnDestroy {

  interval: number;

  angle = 0;

  ngOnInit() {
    this.interval = window.setInterval(() => {
      this.angle = (this.angle + 0.1) % 1;
    }, 50);
  }

  ngOnDestroy() {
    window.clearInterval(this.interval);
  }

  get rotate() {
    return `rotate(${this.angle}turn)`;
  }
}
