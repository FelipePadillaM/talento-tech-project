import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

export interface FileItem {
  id: number;
  name: string;
  date: string;
  isPublic: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class FilesService {
  private files: FileItem[] = [
    { id: 1, name: 'documento.pdf', date: '2024-06-01', isPublic: true },
    { id: 2, name: 'foto.png', date: '2024-06-02', isPublic: false },
    { id: 3, name: 'notas.txt', date: '2024-06-03', isPublic: true }
  ];

  getFiles(): Observable<FileItem[]> {
    return of(this.files);
  }

  addFile(file: Omit<FileItem, 'id'>): Observable<FileItem> {
    const newFile: FileItem = { ...file, id: Date.now() };
    this.files.push(newFile);
    return of(newFile);
  }

  deleteFile(id: number): Observable<boolean> {
    this.files = this.files.filter(f => f.id !== id);
    return of(true);
  }

  renameFile(id: number, newName: string): Observable<FileItem | undefined> {
    const file = this.files.find(f => f.id === id);
    if (file) file.name = newName;
    return of(file);
  }

  togglePrivacy(id: number): Observable<FileItem | undefined> {
    const file = this.files.find(f => f.id === id);
    if (file) file.isPublic = !file.isPublic;
    return of(file);
  }
}
