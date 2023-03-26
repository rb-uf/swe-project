import { Component } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';

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
  ) { }

  ngOnInit() {
    this.getAllSubjects();
  }

  handleError(error: HttpErrorResponse) {
    let errorMessage = 'Unknown error!';
    if (error.error instanceof ErrorEvent) {
      // Client-side errors
      errorMessage = `Error: ${error.error.message}`;
    } else {
      // Server-side errors
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    window.alert(errorMessage);
    return throwError(errorMessage);
  }

  getAllSubjects() {
    return this.http.get<any>('http://localhost:3000/get-subjects').pipe(retry(3), catchError(this.handleError)).subscribe(data => {
      console.log(data);
      this.subjects = data;
    })
  }
}