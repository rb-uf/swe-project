import { Component } from '@angular/core';
import { ReviewService, Review } from './review.service';

@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.css']
})
export class ReviewComponent {
  reviews: Review[] = [];

  constructor(private service: ReviewService) {}
  
  ngOnInit() {
    this.reviews = ReviewService.getReviews();
  }

}
