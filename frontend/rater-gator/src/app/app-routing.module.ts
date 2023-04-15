import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddReviewComponent } from './add-review/add-review.component';
import { SubjectComponent } from './subject/subject.component';

const routes: Routes = [
  { path: 'add-review', component: AddReviewComponent },
  { path: 'subjects', component: SubjectComponent },
  //{ path: 'reviews', component: ReviewComponent },
  //{ path: '', redirectTo: '/reviews', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
