import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [
    CommonModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    FormsModule
  ],
  templateUrl: './auth.component.html',
  styleUrl: './auth.component.scss'
})
export class AuthComponent {
  email = '';
  password = '';
  isRegister = false;

  constructor(private router: Router) {}

  toggleMode() {
    this.isRegister = !this.isRegister;
  }

  submit() {
    if (!this.email || !this.password) {
      alert('Completa todos los campos');
      return;
    }
    // Credenciales por defecto
    if (this.email === 'admin@email.com' && this.password === 'admin123') {
      localStorage.setItem('user', JSON.stringify({
        email: this.email,
        name: 'Admin',
        avatarUrl: 'https://i.pravatar.cc/150?u=admin@email.com',
        registeredAt: '2024-01-01',
        lastLogin: new Date().toISOString().slice(0, 10)
      }));
      this.router.navigateByUrl('/');
      return;
    }
    // Simular registro/login y guardar en localStorage
    localStorage.setItem('user', JSON.stringify({
      email: this.email,
      name: this.email.split('@')[0],
      avatarUrl: 'https://i.pravatar.cc/150?u=' + this.email,
      registeredAt: new Date().toISOString().slice(0, 10),
      lastLogin: new Date().toISOString().slice(0, 10)
    }));
    this.router.navigateByUrl('/');
  }
}
