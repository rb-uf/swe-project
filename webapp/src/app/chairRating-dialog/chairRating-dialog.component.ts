import {Component, Inject, OnDestroy} from "@angular/core";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {ChairRatingsService} from "../chairRatings.service";
import {ChairRating} from "../chairRating";
import {Subscription} from 'rxjs';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {InstantErrorStateMatcher} from "./instant-error-state.matcher";

@Component({
  selector: 'chairRating-dialog',
  templateUrl: 'chairRating-dialog.component.html',
  styleUrls: ['chairRating-dialog.component.scss']
})
export class ChairRatingDialog implements OnDestroy {
  controlGroup: FormGroup;
  errorStateMatcher = new InstantErrorStateMatcher();
  addSubscription: Subscription;
  updateSubscription: Subscription;
  deleteSubscription: Subscription;

  constructor(
    @Inject(MAT_DIALOG_DATA) public chairRating: ChairRating,
    public dialogRef: MatDialogRef<ChairRatingDialog>,
    public service: ChairRatingsService
  ) {
    this.controlGroup = new FormGroup({
      classroomName: new FormControl(chairRating.classroomName, Validators.required),
      rating: new FormControl(chairRating.rating, [
        Validators.required, Validators.min(0), Validators.max(100)
      ])
    });
  }

  save(): void {
    this.chairRating.classroomName = this.formValue('classroomName');
    this.chairRating.rating = this.formValue('rating');
    if (!this.chairRating.id) {
      this.addSubscription = this.service.add(this.chairRating)
        .subscribe(this.dialogRef.close);
    } else {
      this.updateSubscription = this.service.update(this.chairRating)
        .subscribe(this.dialogRef.close);
    }
  }

  delete(): void {
    this.deleteSubscription = this.service.delete(this.chairRating.id)
      .subscribe(this.dialogRef.close);
  }

  hasError(controlName: string, errorCode: string): boolean {
    return !this.controlGroup.valid && this.controlGroup.hasError(errorCode, [controlName]);
  }

  ngOnDestroy(): void {
    if (this.addSubscription) {
      this.addSubscription.unsubscribe();
    }
    if (this.updateSubscription) {
      this.updateSubscription.unsubscribe();
    }
    if (this.deleteSubscription) {
      this.deleteSubscription.unsubscribe();
    }
  }

  private formValue(controlName: string): any {
    return this.controlGroup.get(controlName).value;
  }
}
