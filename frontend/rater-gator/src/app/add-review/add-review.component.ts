import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { AddReviewService } from '../add-review.service';
import { Review } from '../review';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError, map, tap } from 'rxjs/operators';

@Component({
  selector: 'app-add-review',
  templateUrl: './add-review.component.html',
  styleUrls: ['./add-review.component.css']
})
export class AddReviewComponent {

  newReview: Review = {};

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  reviewForm = this.formBuilder.group({
    Subject: '',
    Rating: '',
    Text: '',
    Author: '',
  });

  constructor(
    private formBuilder: FormBuilder,
    private addReviewService: AddReviewService,
    private http: HttpClient,
  ) {}

  onSubmit(): void {
    this.newReview = {
      Subject: <string>this.reviewForm.value.Subject,
      Rating: +<string>this.reviewForm.value.Rating,
      Text: <string>this.reviewForm.value.Text,
      Author: <string>this.reviewForm.value.Author,
    }
    this.addReview();
    console.warn('Your review has been submitted', this.newReview);
    this.newReview = {};
    this.reviewForm.reset();
  }

  addReview() {
    return this.http.post<Review>('http://localhost:3000/create-review', this.newReview, this.httpOptions);
  }

}
