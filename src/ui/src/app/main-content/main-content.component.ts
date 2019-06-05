import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RouteCompatibility, RouteInstallation, RouteVm } from '../shared/shared.const';
import { Keycloak, KeycloakService } from 'keycloak-angular';

@Component({
  selector: 'app-main-content',
  templateUrl: './main-content.component.html',
  styleUrls: ['./main-content.component.less']
})
export class MainContentComponent implements OnInit {
  userDetails: Keycloak.KeycloakProfile;

  constructor(private router: Router,
              private keycloakService: KeycloakService) {
  }

  async ngOnInit() {
    if (await this.keycloakService.isLoggedIn()) {
      this.keycloakService.loadUserProfile().then(res => this.userDetails = res);
      this.router.navigate([`/${RouteCompatibility}/${RouteInstallation}`]).then();
    }
  }

  async doLogout() {
    await this.keycloakService.logout();
  }
}
