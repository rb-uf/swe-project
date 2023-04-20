import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { FormBuilder, FormsModule, ReactiveFormsModule} from '@angular/forms';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { Subject } from './subject/subject.component';

// review object, has 4 properties of user-entered info
export interface Review {
  'Subject': string,
  'Rating': number,
  'Text': string,
  'Author': string,
  'ID'?: number,
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

  postReview(review: Review)  {
    this.http.post<any>('http://localhost:3000/create-review', review).subscribe(data => {
      console.log("New review inserted: ", data);
    });
  }
  
}