import { Injectable } from '@angular/core';

export class Review {
  title: string;
  rating: number; // TODO: restrict range from 0 to 5
  building: string; // TODO: make this an enum of acceptable buildings on campus
  //description: string;
  reviewer: string;

  constructor(title: string, rating: number, building: string, reviewer: string) {
    this.title = title;
    this.rating = rating;
    this.building = building;
    this.reviewer = reviewer;
  }
}

@Injectable({
  providedIn: 'root'
})
export class ReviewsService {
  private _reviews : Review[] = [];

  constructor() { }

  add(review: Review): Review[] {
    this._reviews.push(review);
    return this._reviews;
  }

  get(): readonly Review[] {
    return this._reviews;
  }
}
