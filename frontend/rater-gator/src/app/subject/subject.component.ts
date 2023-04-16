import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { FormBuilder } from '@angular/forms';
import { Subject } from '../subject';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-subject',
  templateUrl: './subject.component.html',
  styleUrls: ['./subject.component.css']
})
export class SubjectComponent {

  subjects: Subject[] = [];
  newSubject: Subject = {};

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
  ) { }

  subjectForm = this.fb.group({
    Name: '',
  })

    // THIS CURRENTLY DOESN'T WORK !!!!!!!!!!!!!!!!
  onSubmit(): void {
    this.newSubject = {
      Name: <string>this.subjectForm.value.Name,
    }
    console.log(this.addSubject());
    this.newSubject = {};
    this.subjectForm.reset();
  }

  // THIS DOESN'T WORK ANYMORE EITHER IDK WHY :)))))))))
  addSubject(): any {
    return this.http.post<any>('http://localhost:3000/create-subject', this.newSubject, this.httpOptions).subscribe(data => {
      console.log(data);
    });;
  }
}