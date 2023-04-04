import { Component } from '@angular/core';
import { Review } from '../review';
import { REVIEWS } from '../mock-reviews';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.css']
})
export class ReviewsComponent {
  reviews: Review[] = REVIEWS;

  selectedReview?: Review;
  onSelect(review: Review): void {
    this.selectedReview = review;
  }
}
