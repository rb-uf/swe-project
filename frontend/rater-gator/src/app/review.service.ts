import { Injectable } from '@angular/core';
import { REVIEWS } from './mock-reviews';
import { Review } from './review';
import { Observable, of } from 'rxjs';
import { MessageService } from './message.service';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  constructor(private http: HttpClient,
    private messageService: MessageService,
  ) { }

  getReviews(): Observable<Review[]> {
    return of(REVIEWS);
  }

  getReview(id: number): Observable<Review> {
    // For now, assume that a hero with the specified `id` always exists.
    // Error handling will be added in the next step of the tutorial.
    const review = REVIEWS.find(h => h.ID === id)!;
    this.messageService.add(`ReviewService: fetched review id=${id}`);
    return of(review);
  }

  /** Log a HeroService message with the MessageService */
  private log(message: string) {
    this.messageService.add(`ReviewService: ${message}`);
  }

}