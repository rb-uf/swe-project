import { Component } from '@angular/core';
import { Review, ReviewsService } from './reviews.service';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.css'],
})
export class ReviewsComponent {
  reviews: readonly Review[] = [];

  constructor(private reviewsService: ReviewsService) {
    this.reviews = reviewsService.get();
  }

  addReview(title: string, rating: number, building: string, reviewer: string) {
    this.reviews = this.reviewsService.add(new Review(title, rating, building, reviewer));
  }
}
