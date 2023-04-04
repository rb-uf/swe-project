import { Component } from '@angular/core';
import { Review } from '../review';

@Component({
  selector: 'app-reviews',
  templateUrl: './reviews.component.html',
  styleUrls: ['./reviews.component.css']
})
export class ReviewsComponent {
  review: Review = {
    ID: 1234,
    Subject: "Newell Hall",
    Rating: 5,
    Text: "very cool",
    Author: "Shane",
    AuthorID: 42069,
  }
}
