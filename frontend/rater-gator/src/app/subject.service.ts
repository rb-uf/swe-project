import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Subject } from './subject';
import { Review } from './review';

@Injectable({
  providedIn: 'root'
})
export class SubjectService {

  constructor(private http: HttpClient) { }

  getSubjects(): Observable<Subject[]> {
    return this.http.get<Subject[]>('/get-subjects');
  }

  getReviews(name: Subject): Observable<Review[]> {
    let params = new HttpParams();
    params = params.append("Name", name.Name);
    return this.http.get<Review[]>('get-subject-reviews', {params:params});
  }
}
