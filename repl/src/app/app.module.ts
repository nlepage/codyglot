import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AceEditorModule } from 'ng2-ace-editor';

import { AppComponent } from './app.component';
import { LoadingComponent } from './loading/loading.component';
import { SourceComponent } from './source/source.component';
import { StdxxxComponent } from './stdxxx/stdxxx.component';
import { ToolbarComponent } from './toolbar/toolbar.component';

import { ExecuteService } from './execute.service';
import { LanguagesService } from './languages.service';

@NgModule({
  bootstrap: [AppComponent],
  declarations: [
    AppComponent,
    ToolbarComponent,
    SourceComponent,
    StdxxxComponent,
    LoadingComponent,
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
})
export class AppModule {}
