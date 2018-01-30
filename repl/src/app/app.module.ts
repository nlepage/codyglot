import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AceEditorModule } from 'ng2-ace-editor';

import { AppComponent } from './app.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { SourceComponent } from './source/source.component';
import { StdxxxComponent } from './stdxxx/stdxxx.component';


@NgModule({
  declarations: [
    AppComponent,
    ToolbarComponent,
    SourceComponent,
    StdxxxComponent
  ],
  imports: [
    BrowserModule,
    AceEditorModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
