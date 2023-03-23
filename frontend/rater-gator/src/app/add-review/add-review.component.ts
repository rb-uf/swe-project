import { Component } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { DeleteComponent } from '../delete/delete.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-review',
  templateUrl: './add-review.component.html',
  styleUrls: ['./add-review.component.css']
})
export class AddReviewComponent {
  public breakpoint: number; // Breakpoint observer code
  public location: string = ``;
  public rating: number = 0;
  public description: string = ``;
  public author: string = ``;
  public addReviewForm: FormGroup;
  wasFormChanged = false;

  constructor(
    private fb: FormBuilder,
    public dialog: MatDialog
  ) { }

  public ngOnInit(): void {
    this.addReviewForm = this.fb.group({
      IdProof: null,
      location: [this.location, [Validators.required, Validators.pattern('[a-zA-Z]+([a-zA-Z ]+)*')]]
    });
    this.breakpoint = window.innerWidth <= 600 ? 1 : 2; // Breakpoint observer code
  }

  public onAddReview(): void {
    this.markAsDirty(this.addReviewForm);
  }

  openDialog(): void {
    console.log(this.wasFormChanged);
    if (this.addReviewForm.dirty) {
      const dialogRef = this.dialog.open(DeleteComponent, {
        width: '540px',
      });
    } else {
      this.dialog.closeAll();
    }
  }

  // tslint:disable-next-line:no-any
  public onResize(event: any): void {
    this.breakpoint = event.target.innerWidth <= 600 ? 1 : 2;
  }

  private markAsDirty(group: FormGroup): void {
    group.markAsDirty();
    // tslint:disable-next-line:forin
    for (const i in group.controls) {
      group.controls[i].markAsDirty();
    }
  }

  formChanged() {
    this.wasFormChanged = true;
  }
}
