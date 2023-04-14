import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { AddReviewService } from '../add-review.service';

@Component({
  selector: 'app-add-review',
  templateUrl: './add-review.component.html',
  styleUrls: ['./add-review.component.css']
})
export class AddReviewComponent {

  newReview = this.addReviewService.getNewReview();

  reviewForm = this.formBuilder.group({
    Subject: '',
    Rating: '',
    Text: '',
    Author: '',
  });

  constructor(
    private formBuilder: FormBuilder,
    private addReviewService: AddReviewService,
  ) {}

  onSubmit(): void {
    this.newReview = this.addReviewService.clear();
    console.warn('Your review has been submitted', this.reviewForm.value);
    this.reviewForm.reset();
  }

}
