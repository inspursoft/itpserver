import { HTTP_INTERCEPTORS, HttpErrorResponse, HttpInterceptor, HttpRequest, HttpResponse } from '@angular/common/http';
import { HttpHandler } from '@angular/common/http/src/backend';
import { HttpEvent } from '@angular/common/http/src/response';
import { Injectable } from '@angular/core';
import { Observable, of, throwError, TimeoutError } from 'rxjs';
import { catchError, map, switchMap, tap, timeout } from 'rxjs/operators';
import { NzNotificationService } from 'ng-zorro-antd';
import { KeycloakService } from 'keycloak-angular';
import { fromPromise } from 'rxjs/internal-compatibility';

@Injectable()
export class HttpClientInterceptor implements HttpInterceptor {

  constructor(private notificationService: NzNotificationService,
              private keycloakService: KeycloakService) {

  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    let authReq: HttpRequest<any> = req.clone({
      headers: req.headers
    });
    const obsToken = fromPromise(this.keycloakService.getToken());
    return obsToken.pipe(switchMap(res => {
      authReq = authReq.clone({
        headers: authReq.headers.set('Authorization', res)
      });
      return next.handle(authReq)
        .pipe(
          timeout(300000),
          catchError((err: HttpErrorResponse | TimeoutError) => {
            if (err instanceof HttpErrorResponse) {
              if (err.status >= 200 && err.status < 300) {
                const res = new HttpResponse({
                  body: null,
                  headers: err.headers,
                  status: err.status,
                  statusText: err.statusText,
                  url: err.url
                });
                return of(res);
              } else if (err.status === 502) {
                this.notificationService.error(`502错误`, err.message);
              } else if (err.status === 504) {
                this.notificationService.error(`504错误`, err.message);
              } else if (err.status === 500) {
                console.log(err);
                this.notificationService.error(`500错误`, err.message);
              } else if (err.status === 400) {
                this.notificationService.error(`400错误`, err.message);
              } else if (err.status === 401) {
                this.notificationService.error(`401错误`, err.message);
              } else if (err.status === 403) {
                this.notificationService.error(`403错误`, err.message);
              } else if (err.status === 404) {
                this.notificationService.error(`404错误`, err.message);
              } else if (err.status === 412) {
                this.notificationService.error(`412错误`, err.message);
              } else if (err.status === 422) {
                this.notificationService.error(`422错误`, err.message);
              } else {
                this.notificationService.error(`${err.status}错误`, err.message);
              }
            } else {
              this.notificationService.error(`访问超时`, err.message, {nzData: err});
            }
            return throwError(err);
          }));
    }));
  }
}

export const HttpInterceptorService = {
  provide: HTTP_INTERCEPTORS,
  useClass: HttpClientInterceptor,
  deps: [NzNotificationService, KeycloakService],
  multi: true
};
