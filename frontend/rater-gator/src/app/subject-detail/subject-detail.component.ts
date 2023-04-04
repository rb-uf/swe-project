import { Component } from '@angular/core';
import { SubjectService } from '../subject.service';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { ReviewService } from '../review.service';

@Component({
  selector: 'app-subject-detail',
  templateUrl: './subject-detail.component.html',
  styleUrls: ['./subject-detail.component.css']
})
export class SubjectDetailComponent {

  constructor(
    private subjectService: SubjectService,
    private route: ActivatedRoute,
    private location: Location) {}

  getReviews(name: any): void {
    const name = Number(this.route.snapshot.paramMap.get('name'));
    this.subjectService.getReviews(name)
      .subscribe(reviews => this.reviews = reviews);
  }
}
