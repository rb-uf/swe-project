import { Component } from '@angular/core';
import { ReviewService, Review } from './review.service';
import { CreateReviewComponent } from './create-review/create-review.component';
import { SubjectComponent } from './subject/subject.component';

import { FormBuilder, FormsModule, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-review',
  templateUrl: './review.component.html',
  styleUrls: ['./review.component.css']
})
export class ReviewComponent {

  // formBuilder object of text-entry fields
  reviewForm = this.formBuilder.group({
    subject: '',
    rating: '',
    description: '',
    author: '',
  });

  constructor(
    private formBuilder: FormBuilder,
    private service: ReviewService,
  ) {}

  // method called to submit review
  // event binded to "Submit Review" button click
  onSubmit(): void {
    let newReview: Review = {
      'Subject': <string>this.reviewForm.value.subject,
      'Rating': +<string>this.reviewForm.value.rating,
      'Text': <string>this.reviewForm.value.description,
      'Author': <string>this.reviewForm.value.author,
    }
    this.service.postReview(newReview); //response returned here
    this.reviewForm.reset();
  }
}
