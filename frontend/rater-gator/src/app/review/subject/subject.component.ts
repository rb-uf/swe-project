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
  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
  ) { }

  ngOnInit() {
    this.getAllSubjects();

  }

  getAllSubjects() {
    return this.http.get<any>('http://localhost:3000/get-subjects').subscribe(data => {
      //console.log(data);
      this.subjects = data;
    })
  }

  getSelectedReviews(subjects: string[]) {
    let selectedSubjects: Subject[] = []
    subjects.forEach(function(item){
      let selectedSubject: Subject = {Name: item}
      selectedSubjects.push(selectedSubject)

    });
    // console.log(selectedSubjects);
    selectedSubjects.forEach((item) =>{
      this.sendGetReviewsRequest(item).subscribe((data: Review[]) => {
        console.log(data);
        data.forEach((subItem: Review) => {
          this.reviews.push(subItem);
        });
      })
    });
    // console.log(this.reviews);
  }

  sendGetReviewsRequest(subject: Subject): Observable<Review[]> {
    const options = subject ? 
    {params : new HttpParams().set('Name', subject.Name)} : {};

    // console.log( this.http.post<any>(`http://localhost:3000/get-subject-reviews`, options));

    return this.http.post<Review[]>(`http://localhost:3000/get-subject-reviews`, subject);
  }

  subjectForm = this.fb.group({
    Name: '',
  })

  onSubmit(): void {
    let newSubject = {
      Name: <string>this.subjectForm.value.Name,
    }
    //console.log(newSubject);
    this.addSubject(newSubject); //response returned here
    this.subjectForm.reset();
    //this.loadReviews();
  }

  addSubject(newSubject: Subject): any {
    this.http.post<any>('http://localhost:3000/create-subject', newSubject).subscribe(data => {
      //console.log(data);
    });
    this.getAllSubjects();
  }

  onChange(selectedOptions: string[]) {
    //console.log(selectedOptions);
    this.getSelectedReviews(selectedOptions);
  }
}