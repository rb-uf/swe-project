import { Component } from '@angular/core';
import { Subject } from '../subject';
import { SubjectService } from '../subject.service';

@Component({
  selector: 'app-subjects',
  templateUrl: './subjects.component.html',
  styleUrls: ['./subjects.component.css']
})
export class SubjectsComponent {
  subjects: Subject[] = [];

  constructor(private subjectService: SubjectService) {}

  getSubjects(): void {
    this.subjectService.getSubjects()
        .subscribe(subjects => this.subjects = subjects);
  }

  ngOnInit(): void {
    this.getSubjects();
  }
}
