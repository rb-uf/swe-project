import { Injectable } from '@angular/core';
import { REVIEWS } from './mock-reviews';
import { Review } from './review';
import { Observable, of } from 'rxjs';
import { MessageService } from './message.service';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {

  constructor(private messageService: MessageService) { }

  getReviews(): Observable<Review[]> {
    const reviews = of(REVIEWS);
    this.messageService.add('ReviewService: fetched reviews');
    return reviews;
  }
}