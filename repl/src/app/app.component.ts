import { Component } from '@angular/core';
import { LanguageInfo } from './language.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  language: LanguageInfo;
  source: string;
  stdin: string;
  stdout: string = 'stdout\nstdout';
  stderr: string = 'stderr\nstderr';
}
