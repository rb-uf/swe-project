import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Subject } from './subject';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SubjectService {

  constructor(
    private http: HttpClient,
  ) { }

  sendPostRequest(newSubject: Subject): Observable<Subject> {
    return this.http.post<Subject>('http://localhost:3000/create-subject', newSubject);
  }
}
