import { Injectable } from '@angular/core';

export interface Review {
  title: string;
  rating: number; // TODO: restrict range from 0 to 5
  building: string; // TODO: make this an enum of acceptable buildings on campus
  //description: string;
  reviewer: string;
}

@Injectable({
  providedIn: 'root'
})
export class ReviewsService {

  constructor() { }

  get(): Review[] {
    return [
      {
        title: "Uncomfortable and Squeaky",
        rating: 2,
        building: "Carleton Auditorium",
        reviewer: "ShaneF"
      },
      {
        title: "Nicest chairs ever",
        rating: 5,
        building: "Marston Science Library",
        reviewer: "DevalaG"
      }
    ];
  }
}
