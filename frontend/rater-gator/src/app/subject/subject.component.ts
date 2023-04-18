import { Component } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { FormBuilder } from '@angular/forms';
import { Subject } from '../subject';
import { Observable } from 'rxjs';
import { SubjectService } from '../subject.service';

@Component({
  selector: 'app-subject',
  templateUrl: './subject.component.html',
  styleUrls: ['./subject.component.css']
})
export class SubjectComponent {

  subjects: Subject[] = [];
  newSubject: Subject = {};

  constructor(
    private fb: FormBuilder,
    private subjectService: SubjectService,
  ) { }

  subjectForm = this.fb.group({
    Name: '',
  })

    // THIS CURRENTLY DOESN'T WORK !!!!!!!!!!!!!!!!
  onSubmit(): void {
    this.newSubject = {
      Name: <string>this.subjectForm.value.Name,
    }
    this.subjectService.sendPostRequest(this.newSubject).subscribe(
      res => {
        console.log(res);
      }
    );
    this.newSubject = {};
    this.subjectForm.reset();
  }
}