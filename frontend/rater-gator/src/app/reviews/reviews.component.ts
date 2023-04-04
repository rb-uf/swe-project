import { Component, OnInit } from '@angular/core';
import { Review } from '../review';
import { ReviewService } from '../review.service';
import { MessageService } from '../message.service';


@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.css']
})
export class ReviewsComponent {
  reviews: Review[] = [];
  selectedReview?: Review;

  constructor(private reviewService: ReviewService, private messageService: MessageService) {}

  getReviews(): void {
    this.reviewService.getReviews()
        .subscribe(reviews => this.reviews = reviews);
  }

  
  onSelect(review: Review): void {
    this.selectedReview = review;
  }

  ngOnInit(): void {
    this.getReviews();
  }

}
