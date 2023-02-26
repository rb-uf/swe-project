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
    this.loadReviews();
  }

  loadReviews() {
    this.reviews = this.service.getReviews();
  }

  // method called to submit review
  // event binded to "Submit Review" button click
  onSubmit(): void {
    let newReview = {
      location: <string>this.reviewForm.value.location,
      rating: +<string>this.reviewForm.value.rating,
      description: <string>this.reviewForm.value.description,
      author: <string>this.reviewForm.value.author,
    }

    this.service.addNewReview(newReview) //response returned here
      .subscribe(); 
    this.reviewForm.reset();
    this.loadReviews();
  }
}
