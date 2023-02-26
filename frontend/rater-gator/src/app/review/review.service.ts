import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { FormBuilder, FormsModule, ReactiveFormsModule, FormGroup, FormControl } from '@angular/forms';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

// review object, has 4 properties of user-entered info
export interface Review {
  location: string;
  rating: number;
  description: string;
  author: string;
}

@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  // service list for reviews
  // initalized with one fake review
  _reviews: Review[] = [{
    location: "Marston",
    rating: 5,
    description: "comfy",
    author: "Shane",
  }];

  constructor(
    
    private http: HttpClient,
  ) { }

  // reviews list accessor
  getReviews(): Review[] {
    return this._reviews;
  }

  // new review mutator
  addNewReview(reviewForm: FormGroup<{ location: FormControl<string | null>; rating: FormControl<string | null>; description: FormControl<string | null>; author: FormControl<string | null>; }>) {
    this._reviews.push({
      location: <string>reviewForm.value.location,
      rating: +<string>reviewForm.value.rating,
      description: <string>reviewForm.value.description,
      author: <string>reviewForm.value.author,
    });
    console.warn('Your review has been submitted', reviewForm.value);
  }

  postReview(review: Review): Observable<Review> {
    return this.http.post<Review>('/create-subject', review);
  }
}