import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule }   from '@angular/forms';
import { AceEditorModule } from 'ng2-ace-editor';

import { AppComponent } from './app.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { SourceComponent } from './source/source.component';
import { StdxxxComponent } from './stdxxx/stdxxx.component';

import { LanguagesService } from './languages.service';
import { ExecuteService } from './execute.service';

@NgModule({
  declarations: [
    AppComponent,
    ToolbarComponent,
    SourceComponent,
    StdxxxComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    AceEditorModule,
  ],
  providers: [
    LanguagesService,
    ExecuteService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
