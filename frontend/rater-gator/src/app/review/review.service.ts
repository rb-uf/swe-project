import { Injectable } from '@angular/core';

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

  constructor() { }

  static getReviews(): Review[] {
    return [{
      location: "Marston",
      rating: 5,
      description: "comfy",
      author: "Shane",
    },
    {
      location: "Carleton",
      rating:25,
      description: "uncomfy",
      author: "Devala",
    }]
  }
}
