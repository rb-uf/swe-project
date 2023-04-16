import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { Review } from '../review';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Subject } from '../subject';

@Component({
  selector: 'app-add-review',
  templateUrl: './add-review.component.html',
  styleUrls: ['./add-review.component.css']
})
export class AddReviewComponent {

  newReview: Review = {};
  subjects: Subject[] = [];

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  reviewForm = this.formBuilder.group({
    Subject: '',
    Rating: '',
    Text: '',
    Author: '',
  });

  constructor(
    private formBuilder: FormBuilder,
    private http: HttpClient,
  ) {}

  ngOnInit() {
    this.getAllSubjects();
  }

  getAllSubjects() {
    return this.http.get<Subject[]>('http://localhost:3000/get-subjects').subscribe(data => {
      console.log(data);
      this.subjects = data;
    })
  }

  onSubmit(): void {
    this.newReview = {
      Subject: <string>this.reviewForm.value.Subject,
      Rating: +<string>this.reviewForm.value.Rating,
      Text: <string>this.reviewForm.value.Text,
      Author: <string>this.reviewForm.value.Author,
    }
    console.log(this.addReview());
    console.warn('Your review has been submitted', this.newReview);
    this.newReview = {};
    this.reviewForm.reset();
  }

  onChange(selectedSubject: string) {
    console.log(selectedSubject);
  }

  // this works???? wtf
  addReview() {
    return this.http.post<Review>('http://localhost:3000/create-review', this.newReview, this.httpOptions);
  }

}
