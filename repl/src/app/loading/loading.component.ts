import { Component, OnDestroy, OnInit } from "@angular/core";

@Component({
  selector: "loading",
  templateUrl: "./loading.component.html",
})
export class LoadingComponent implements OnInit, OnDestroy {

  public interval: number;

  public angle = 0;

  public ngOnInit() {
    this.interval = window.setInterval(() => {
      this.angle = (this.angle + 0.1) % 1;
    }, 50);
  }

  public ngOnDestroy() {
    window.clearInterval(this.interval);
  }

  get rotate() {
    return `rotate(${this.angle}turn)`;
  }
}
