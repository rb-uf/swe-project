import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { Review, ReviewService } from '../review/review.service';

@Component({
  selector: 'app-create-review',
  templateUrl: './create-review.component.html',
  styleUrls: ['./create-review.component.css']
})

export class CreateReviewComponent {

  // list of new reviews enetered by user
  static newReviews: Review[] = [];

  // formBuilder object of text-entry fields
  newReviewForm = this.formBuilder.group({
    location: '',
    rating: '',
    description: '',
    author: '',
  });

  constructor (
    private formBuilder: FormBuilder,
  ) {}

  // method called to submit review
  // event binded to "Submit Review" button click
  onSubmit(): void {
    CreateReviewComponent.newReviews = [{
      location: <string>this.newReviewForm.value.location,
      rating: +<string>this.newReviewForm.value.rating,
      description: <string>this.newReviewForm.value.description,
      author: <string>this.newReviewForm.value.author,
    }]
    ReviewService.getNewReviews();
    console.warn('Your review has been submitted', this.newReviewForm.value);
    this.newReviewForm.reset();
  }
}
