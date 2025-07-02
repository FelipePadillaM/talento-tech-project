import { Component } from '@angular/core';
import { RouterOutlet, RouterModule } from '@angular/router';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { UserService, User } from './services/user.service';
import { Observable } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { UserInfoDialogComponent } from './components/user-info/user-info-dialog.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    RouterModule,
    MatToolbarModule,
    MatButtonModule,
    MatMenuModule,
    MatIconModule,
    FormsModule,
    CommonModule
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'dashboard-app';
  user$!: Observable<User>;

  constructor(private userService: UserService, private dialog: MatDialog) {}

  ngOnInit() {
    this.user$ = this.userService.getUser();
  }

  showInfo(user: User) {
    this.dialog.open(UserInfoDialogComponent, {
      data: {
        name: user.name,
        email: user.email,
        avatarUrl: user.avatarUrl,
        registeredAt: '2024-01-01', // mock
        lastLogin: '2024-07-01' // mock
      },
      width: '400px'
    });
  }

  logout() {
    localStorage.removeItem('user');
    window.location.href = '/auth';
  }
}
