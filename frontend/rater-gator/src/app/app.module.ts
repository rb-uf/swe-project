import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { ReviewComponent } from './review/review.component';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SubmitButtonComponent } from './review/create-review/submit-button/submit-button.component';
import { CreateReviewComponent } from './review/create-review/create-review.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { SubjectComponent } from './review/subject/subject.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { MatFormFieldModule } from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';


@NgModule({
  declarations: [
    AppComponent,
    ReviewComponent,
    SubmitButtonComponent,
    CreateReviewComponent,
    SubjectComponent,
    ToolbarComponent,
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    MatFormFieldModule,
    MatSelectModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
