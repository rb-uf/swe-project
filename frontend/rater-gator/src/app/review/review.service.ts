import { Injectable } from '@angular/core';

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
  _reviews: Review[] = [{
    location: "Marston",
    rating: 5,
    description: "comfy",
    author: "Shane",
  }];

  // list of new reviews enetered by user
  _newReviews: Review[] = [];

  constructor() { }

  // reviews list accessor
  getReviews(): Review[] {
    return this._reviews;
  }

  // new reviews accessor
  getNewReviews() {
    this._reviews.push(this._newReviews[0]);
  }
}