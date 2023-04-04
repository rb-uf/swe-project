import { Component, Input } from '@angular/core';
import { Review } from '../review';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { ReviewService } from '../review.service';


@Component({
  selector: 'app-review-detail',
  templateUrl: './review-detail.component.html',
  styleUrls: ['./review-detail.component.css']
})
export class ReviewDetailComponent {
  @Input() review?: Review;

  constructor(
    private route: ActivatedRoute,
    private reviewService: ReviewService,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.getReview();
  }
  
  getReview(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    this.reviewService.getReview(id)
      .subscribe(review => this.review = review);
  }

  goBack(): void {
    this.location.back();
  }
}
