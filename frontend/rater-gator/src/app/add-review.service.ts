import { Injectable } from '@angular/core';
import { Review } from './review';

@Injectable({
  providedIn: 'root'
})
export class AddReviewService {

  newReviews: Review[] = [];

  constructor() { }

  getNewReview() {
    return this.newReviews;
  }

  clear() {
    this.newReviews = [];
    return this.newReviews;
  }
}
