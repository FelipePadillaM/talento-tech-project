import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { FilesService, FileItem } from './files.service';
import { UserService, User } from '../../services/user.service';
import { Observable } from 'rxjs';

interface EditingState {
  [id: number]: { editing: boolean; newName: string };
}

@Component({
  selector: 'app-files',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule
  ],
  templateUrl: './files.component.html',
  styleUrl: './files.component.scss'
})
export class FilesComponent implements OnInit {
  files$!: Observable<FileItem[]>;
  editingState: EditingState = {};
  selectedFile: File | null = null;
  displayedColumns = ['name', 'date', 'privacy', 'actions'];
  user$!: Observable<User>;

  constructor(private filesService: FilesService, private userService: UserService) {}

  ngOnInit(): void {
    this.loadFiles();
    this.user$ = this.userService.getUser();
  }

  loadFiles() {
    this.files$ = this.filesService.getFiles();
  }

  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
    } else {
      this.selectedFile = null;
    }
  }

  uploadFile() {
    if (this.selectedFile) {
      const newFile = {
        name: this.selectedFile.name,
        date: new Date().toISOString().slice(0, 10),
        isPublic: false
      };
      this.filesService.addFile(newFile).subscribe(() => {
        this.loadFiles();
        this.selectedFile = null;
      });
    }
  }

  deleteFile(id: number) {
    this.filesService.deleteFile(id).subscribe(() => this.loadFiles());
  }

  startEditing(file: FileItem) {
    this.editingState[file.id] = { editing: true, newName: file.name };
  }

  cancelEditing(file: FileItem) {
    this.editingState[file.id] = { editing: false, newName: file.name };
  }

  renameFile(id: number) {
    const state = this.editingState[id];
    if (state) {
      this.filesService.renameFile(id, state.newName).subscribe(() => {
        this.loadFiles();
        this.editingState[id] = { editing: false, newName: state.newName };
      });
    }
  }

  togglePrivacy(id: number) {
    this.filesService.togglePrivacy(id).subscribe(() => this.loadFiles());
  }
} 