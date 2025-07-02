import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';

interface PublicFile {
  id: number;
  name: string;
  date: string;
  owner: string;
}

@Component({
  selector: 'app-public-files',
  standalone: true,
  imports: [CommonModule, MatTableModule, MatIconModule],
  templateUrl: './public-files.component.html',
  styleUrl: './public-files.component.scss'
})
export class PublicFilesComponent {
  displayedColumns = ['name', 'date', 'owner', 'actions'];
  publicFiles: PublicFile[] = [
    { id: 1, name: 'manual.pdf', date: '2024-06-01', owner: 'Ana Gómez' },
    { id: 2, name: 'foto-viaje.jpg', date: '2024-06-02', owner: 'Luis Torres' },
    { id: 3, name: 'recetas.docx', date: '2024-06-03', owner: 'Maria Ruiz' }
  ];

  download(file: PublicFile) {
    // Simulación de descarga
    alert(`Descargando: ${file.name}`);
  }
}
