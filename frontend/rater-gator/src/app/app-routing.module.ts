import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ReviewComponent } from './review/review.component';
import { SubjectComponent } from './review/subject/subject.component';

const routes: Routes = [
  { path: 'add-review', component: ReviewComponent },
  { path: '', component: SubjectComponent },
  { path: '', redirectTo: '/', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }