import { Routes } from '@angular/router';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { FilesComponent } from './components/files/files.component';
import { PublicFilesComponent } from './components/public-files/public-files.component';
import { AuthComponent } from './components/auth/auth.component';
import { inject } from '@angular/core';
import { Router } from '@angular/router';

function authGuard() {
  const router = inject(Router);
  const user = localStorage.getItem('user');
  if (!user) {
    router.navigateByUrl('/auth');
    return false;
  }
  return true;
}

export const routes: Routes = [
  { path: 'auth', component: AuthComponent },
  { path: '', component: DashboardComponent, canActivate: [authGuard] },
  { path: 'files', component: FilesComponent, canActivate: [authGuard] },
  { path: 'public-files', component: PublicFilesComponent, canActivate: [authGuard] },
];
