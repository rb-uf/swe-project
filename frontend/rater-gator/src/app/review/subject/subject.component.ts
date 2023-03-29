import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder } from '@angular/forms';

export interface Subject {
  Name: string;
}

@Component({
  selector: 'app-subject',
  templateUrl: './subject.component.html',
  styleUrls: ['./subject.component.css']
})
export class SubjectComponent {

  subjects = [];

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
  ) { }

  ngOnInit() {
    this.getAllSubjects();
  }

  getAllSubjects() {
    return this.http.get<any>('http://localhost:3000/get-subjects').subscribe(data => {
      console.log(data);
      this.subjects = data;
    })
  }

  getSelectedReviews(subjects: string[]) {
    var subjectsObject: Subject = {
      Name: subjects[0], //todo: needs to send all subjects
    }
    console.log(subjects);
    return this.http.get<any>(`http://localhost:3000/get-subject-reviews/${subjectsObject.Name}`).subscribe(data => { 
      console.log(data);
      this.subjects = data;
    })
  }

  subjectForm = this.fb.group({
    Name: '',
  })

  onSubmit(): void {
    let newSubject = {
      Name: <string>this.subjectForm.value.Name,
    }
    console.log(newSubject);
    this.addSubject(newSubject); //response returned here
    this.subjectForm.reset();
    //this.loadReviews();
  }

  addSubject(newSubject: Subject): any {
    this.http.post<any>('http://localhost:3000/create-subject', newSubject).subscribe(data => {
      console.log(data);
    });
    this.getAllSubjects();
  }

  onChange(selectedOptions: string[]) {
    console.log(selectedOptions);
    this.getSelectedReviews(selectedOptions);
  }
}