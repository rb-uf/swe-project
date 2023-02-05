import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";
import {ChairRating} from "./chairRating";

@Injectable({
  providedIn: 'root'
})
export class ChairRatingsService {
  baseUrl: string = '/chairRatings';
  readonly headers = new HttpHeaders()
    .set('Content-Type', 'application/json');

  constructor(private http: HttpClient) {}

  getAll(): Observable<ChairRating[]> {
    return this.http.get<ChairRating[]>(this.baseUrl);
  }

  add(st: ChairRating): Observable<ChairRating> {
    return this.http.post<ChairRating>(this.baseUrl, st, {headers: this.headers});
  }

  update(st: ChairRating): Observable<ChairRating> {
    return this.http.put<ChairRating>(
      `${this.baseUrl}/${st.id}`, st, {headers: this.headers}
    );
  }

  delete(id: string): Observable<ChairRating> {
    return this.http.delete<ChairRating>(`${this.baseUrl}/${id}`);
  }
}
