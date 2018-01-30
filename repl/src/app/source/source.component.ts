import { Component, OnInit, Input } from '@angular/core';
import { LanguageInfo } from '../language.service'
import 'brace';
import 'brace/theme/chrome';
import 'brace/mode/golang';
import 'brace/mode/javascript';

@Component({
  selector: 'app-source',
  templateUrl: './source.component.html',
  styleUrls: ['./source.component.css']
})
export class SourceComponent implements OnInit {

  @Input()
  language: LanguageInfo

  constructor() {}

  ngOnInit() {
  }

}
