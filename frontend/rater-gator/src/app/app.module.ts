import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { ReviewComponent } from './review/review.component';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SubmitButtonComponent } from './review/create-review/submit-button/submit-button.component';
import { CreateReviewComponent } from './review/create-review/create-review.component';

@NgModule({
  declarations: [
    AppComponent,
    ReviewComponent,
    SubmitButtonComponent,
    CreateReviewComponent,
  ],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {}
