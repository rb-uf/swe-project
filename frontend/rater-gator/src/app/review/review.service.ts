import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { FormBuilder, FormsModule, ReactiveFormsModule} from '@angular/forms';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { Subject } from './subject/subject.component';

// review object, has 4 properties of user-entered info
export interface Review {
  subject: string;
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
  _reviews = [];

  constructor(
    
    private http: HttpClient,
  ) { }

  // reviews list accessor
  getReviews(): Review[] {
    let subject: Subject = {
      Name: "Marston:"
    }
    //this.http.get<any>('http://localhost:3000/get-subject-reviews', subject).subscribe(data => {
      //console.log(data);
      //this._reviews = data;
    //})
    return this._reviews;
  }

  postReview(review: Review)  {
    this.http.post<any>('http://localhost:3000/create-review', review).subscribe(data => {
      console.log(data);
    });
    this.getReviews();
  }
  
}