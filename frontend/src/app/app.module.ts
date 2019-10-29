import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FusionMainComponent } from './fusion-main/fusion-main.component';
import { DynamicFormComponent } from './dynamic-form/dynamic-form.component';
import { DynamicFormQuestionComponent } from './dynamic-form-question/dynamic-form-question.component';

@NgModule({
  imports: [ BrowserModule, ReactiveFormsModule, AppRoutingModule ],
  declarations: [ AppComponent, FusionMainComponent, DynamicFormComponent, DynamicFormQuestionComponent  ],
  bootstrap: [ AppComponent ]
})
export class AppModule {
  constructor() {
  }
}
