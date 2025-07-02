import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

export interface User {
  id: number;
  name: string;
  email: string;
  avatarUrl: string;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private user: User = {
    id: 1,
    name: 'Juan Pérez',
    email: 'juan.perez@email.com',
    avatarUrl: 'https://i.pravatar.cc/150?img=3'
  };

  getUser(): Observable<User> {
    return of(this.user);
  }
} 