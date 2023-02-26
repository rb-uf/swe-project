import { Component } from '@angular/core';
import { ReviewService, Review } from './review.service';
import { FormBuilder, FormsModule, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.css']
})
export class ReviewComponent {
  // list of reviews
  reviews: Review[] = [];

  // formBuilder object of text-entry fields
  reviewForm = this.formBuilder.group({
    location: '',
    rating: '',
    description: '',
    author: '',
  });

  constructor(
    private formBuilder: FormBuilder,
    private service: ReviewService,
  ) {}
  
  ngOnInit() {
    this.reviews = this.service.getReviews();
  }

  // method called to submit review
  // event binded to "Submit Review" button click
  onSubmit(): void {
    this.service.addNewReview(this.reviewForm);
    this.reviewForm.reset();
    this.reviews = this.service.getReviews();

    //this.service.postReview(this.service.getNewReview());
  }
}
