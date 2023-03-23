import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import {MatDialogRef,MatDialog} from '@angular/material/dialog';
//import { AddCustomerComponent } from '../add-review/add-review.component';

@Component({
  selector: 'app-delete',
  templateUrl: './delete.component.html',
  styleUrls: ['./delete.component.css']
})
export class DeleteComponent {

  constructor(private fb: FormBuilder,
    private dialog: MatDialog,
                  private dialogRef: MatDialogRef<DeleteComponent>) {} // Closing dialog window
    
    public cancel(): void { // To cancel the dialog window
    this.dialogRef.close();
    }
    
    public cancelN(): void { 
        this.dialog.closeAll();
    }

}
