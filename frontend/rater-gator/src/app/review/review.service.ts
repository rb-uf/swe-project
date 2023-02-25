import { Injectable } from '@angular/core';
import { CreateReviewComponent } from '../create-review/create-review.component';

// review object, has 4 properties of user-entered info
export interface Review {
  location: string;
  rating: number;
  description: string;
  author: string;
}

@Injectable({
  providedIn: 'root'
})
export class ReviewService {
  // service list for reviews
  // initalized with one fake review
  static _reviews: Review[] = [{
    location: "Marston",
    rating: 5,
    description: "comfy",
    author: "Shane",
  // },
  // {
    // location: "Carleton",
    // rating: 2,
    // description: "uncomfy",
    // author: "Devala",
  }];

  constructor() { }

  // reviews list accessor
  static getReviews(): Review[] {
    return this._reviews;
  }

  // new reviews accessor
  static getNewReviews() {
    this._reviews.push(CreateReviewComponent.newReviews[0]);
  }
}