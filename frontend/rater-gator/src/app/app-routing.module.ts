import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ReviewsComponent } from './reviews/reviews.component';
import { HomeComponent } from './home/home.component';
import { ReviewDetailComponent } from './review-detail/review-detail.component';
import { SubjectsComponent } from './subjects/subjects.component';

const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'reviews', component: ReviewsComponent },
  { path: 'subjects', component: SubjectsComponent},
  { path: 'detail/:id', component: ReviewDetailComponent },
  //{ path: 'subject/:name', component: SubjectDetailComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }