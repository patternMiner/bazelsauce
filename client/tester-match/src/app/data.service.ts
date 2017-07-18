import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import 'rxjs/add/operator/map'

@Injectable()
export class DataService {
  private countriesUrl = 'http://localhost:8080/countries';
  private devicesUrl = 'http://localhost:8080/devices';
  private testersUrl = 'http://localhost:8080/tester_match';

  constructor(private http: Http) { }

  public fetchCountries() {
    return this.http.get(this.countriesUrl)
      .map((res: Response) => res.json());
  }

  public fetchDevices() {
    return this.http.get(this.devicesUrl)
      .map((res: Response) => res.json());
  }

  public matchTesters() {
    return this.http.get(this.testersUrl)
      .map((res: Response) => res.json());
  }
}
