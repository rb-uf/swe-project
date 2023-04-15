import { Injectable } from '@angular/core';
import { Review } from './review';

@Injectable({
  providedIn: 'root'
})
export class AddReviewService {

  constructor() { }

  clear() {
    return {};
  }
}
