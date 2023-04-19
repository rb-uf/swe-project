import { Component } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams} from '@angular/common/http';
import { FormBuilder } from '@angular/forms';
import { Review } from '../review.service';
import { Observable } from 'rxjs';

export interface Subject {
  Name: string;
}

@Component({
  selector: 'app-subject',
  templateUrl: './subject.component.html',
  styleUrls: ['./subject.component.css']
})
export class SubjectComponent {

  subjects: Subject[]= [];
  reviews: Review[] = [];
  showEdit: boolean = false;
  editReviewID: number = -1;

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
  ) { }

  ngOnInit() {
    this.getAllSubjects();
  }

  getAllSubjects() {
    return this.http.get<any>('http://localhost:3000/get-subjects').subscribe(data => {
      this.subjects = data;
    })
  }

  getSelectedReviews(subject: string) {
      let selectedSubject: Subject = {Name: subject}

      this.sendGetReviewsRequest(selectedSubject).subscribe((data: Review[]) => {
        data.forEach((subItem: Review) => {
          this.reviews.push(subItem);
        });
      })

    console.log("Reviews: ", this.reviews);
  }

  sendGetReviewsRequest(subject: Subject): Observable<Review[]> {
    let body = {
      'Name': subject.Name,
      'MaxReviews': 10000,
    }

    return this.http.post<Review[]>(`http://localhost:3000/get-subject-reviews`, body);
  }

  subjectForm = this.fb.group({
    Name: '',
  })

  editForm = this.fb.group({
    subject: '',
    rating: '',
    description: '',
    author: '',
  });

  onEditSubmit(): void {
    let editedReview: Review = {
      'ID': <number>this.editReviewID,
      'Subject': <string>this.editForm.value.subject,
      'Rating': +<string>this.editForm.value.rating,
      'Text': <string>this.editForm.value.description,
      'Author': <string>this.editForm.value.author,
    }
    console.log("Edited review: ", editedReview);
    this.editForm.reset();
    this.editReview(editedReview).subscribe(data => {
      console.log(data);
    }); //response returned here
    this.editForm.reset();
  }

  editReview(editedReview: Review) {
    return this.http.put<any>('http://localhost:3000/update-review', editedReview);
  }

  onSubmit(): void {
    let newSubject = {
      Name: <string>this.subjectForm.value.Name,
    }
    this.addSubject(newSubject); //response returned here
    this.subjectForm.reset();
  }

  addSubject(newSubject: Subject): any {
    this.http.post<any>('http://localhost:3000/create-subject', newSubject).subscribe(data => {
    });
    this.getAllSubjects();
  }

  onChange(selectedOption: string) {
    this.reviews = [];
    this.getSelectedReviews(selectedOption);
  }

  onDeleteClick(reviewID: number) { //(click)="onDeleteClick(review.ID)"
    let body = {
      'ID': reviewID,
    }
    //this.http.delete<any>('http://localhost:3000/delete-review', body).subscribe(data => {

    //});
  }

  onEditClick(reviewID: any) { //(click)="onEditClick(review.ID)"
    this.showEdit = !this.showEdit;
    this.editReviewID = reviewID;
  }
}