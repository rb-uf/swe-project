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
  newReviewForm = this.formBuilder.group({
    location: '',
    rating: '',
    description: '',
    author: '',
  });

  constructor(
    private service: ReviewService,
    private formBuilder: FormBuilder,
  ) {}
  
  ngOnInit() {
    this.reviews = this.service.getReviews();
  }

  // method called to submit review
  // event binded to "Submit Review" button click
  onSubmit(): void {
    this.service.newReviews = [{
      location: <string>this.newReviewForm.value.location,
      rating: +<string>this.newReviewForm.value.rating,
      description: <string>this.newReviewForm.value.description,
      author: <string>this.newReviewForm.value.author,
    }]
    this.service.getNewReviews();
    console.warn('Your review has been submitted', this.newReviewForm.value);
    this.newReviewForm.reset();
  }
}
