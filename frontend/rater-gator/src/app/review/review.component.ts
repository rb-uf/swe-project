import { Component } from '@angular/core';
import { ReviewService, Review } from './review.service';
import { CreateReviewComponent } from '../create-review/create-review.component';

@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.css']
})
export class ReviewComponent {
  // list of reviews
  reviews: Review[] = [];

  constructor(private service: ReviewService) {}
  
  ngOnInit() {
    this.reviews = ReviewService.getReviews();
  }

}
