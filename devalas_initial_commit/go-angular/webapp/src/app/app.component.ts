import {Component, OnDestroy, OnInit} from '@angular/core';
import {MatDialog} from "@angular/material/dialog";
import {ChairRatingDialog} from "./chairRating-dialog/chairRating-dialog.component";
import {ChairRatingsService} from "./chairRatings.service";
import {ChairRating} from "./chairRating";
import {Subscription} from "rxjs";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit, OnDestroy {
  displayedColumns = ['classroomName', 'rating'];
  dataSource: ChairRating[] = [];
  getAllSubscription: Subscription;
  dialogSubscription: Subscription;

  constructor(public dialog: MatDialog, public service: ChairRatingsService) {}

  openEditDialog(st: ChairRating) {
    this.openDialog(new ChairRating(st.id, st.classroomName, st.rating));
  }

  openNewDialog(): void {
    this.openDialog(new ChairRating());
  }

  private openDialog(st: ChairRating): void {
    this.dialogSubscription = this.dialog
      .open(ChairRatingDialog, {data: st, minWidth: '30%'})
      .afterClosed().subscribe(() => this.loadChairRatingsList());
  }

  private loadChairRatingsList(): void {
    this.getAllSubscription = this.service.getAll()
      .subscribe(chairRatings => this.dataSource = chairRatings);
  }

  ngOnInit(): void {
    this.loadChairRatingsList();
  }

  ngOnDestroy(): void {
    this.getAllSubscription.unsubscribe();
    if (this.dialogSubscription) {
      this.dialogSubscription.unsubscribe();
    }
  }
}
