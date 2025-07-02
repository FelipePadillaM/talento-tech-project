import { Component, Inject } from '@angular/core';
import { MatDialogModule, MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

export interface UserInfoData {
  name: string;
  email: string;
  avatarUrl: string;
  registeredAt: string;
  lastLogin: string;
}

@Component({
  selector: 'app-user-info-dialog',
  standalone: true,
  imports: [
    CommonModule,
    MatDialogModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule
  ],
  templateUrl: './user-info-dialog.component.html',
  styleUrl: './user-info-dialog.component.scss'
})
export class UserInfoDialogComponent {
  constructor(
    public dialogRef: MatDialogRef<UserInfoDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: UserInfoData
  ) {}

  close(): void {
    this.dialogRef.close();
  }

  save(): void {
    alert('Cambios guardados (simulado)');
    this.dialogRef.close(this.data);
  }
}
