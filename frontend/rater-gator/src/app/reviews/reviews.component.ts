import { Component } from '@angular/core';
import { Review, ReviewsService } from './reviews.service';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.css'],
})
export class ReviewsComponent {
  reviews: Review[] = [];

  constructor(private reviewsService: ReviewsService) {
    this.reviews = reviewsService.get();
  }
}
