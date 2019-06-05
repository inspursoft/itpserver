import { Component, OnInit } from '@angular/core';
import { Keycloak, KeycloakService } from 'keycloak-angular';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less'],
})
export class AppComponent implements OnInit {
  userDetails: Keycloak.KeycloakProfile;

  constructor(private keycloakService: KeycloakService) {}

  async ngOnInit() {
    // if (await this.keycloakService.isLoggedIn()) {
    //   this.userDetails = await this.keycloakService.loadUserProfile();
    // }
  }

  async doLogout() {
    await this.keycloakService.logout();
  }
}
