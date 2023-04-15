import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
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

  constructor(
    private http: HttpClient,
    private fb: FormBuilder,
  ) { }

  subjectForm = this.fb.group({
    Name: '',
  })

    // THIS CURRENTLY DOESN'T WORK !!!!!!!!!!!!!!!!
  onSubmit(): void {
    let newSubject: Subject = {
      Name: <string>this.subjectForm.value.Name,
    }
    console.log(newSubject);
    this.addSubject(newSubject).subscribe(data => {
      console.log(data)
    });; //response returned here
    this.subjectForm.reset();
  }

  // THIS DOESN'T WORK ANYMORE EITHER IDK WHY :)))))))))
  addSubject(newSubject: Subject): Observable<Subject> {
    return this.http.post('http://localhost:3000/create-subject', newSubject);
  }
}