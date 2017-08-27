import { Injectable } from '@angular/core';
import { Http, Response, URLSearchParams } from '@angular/http';
import 'rxjs/add/operator/map'

@Injectable()
export class DataService {
  private countriesUrl = 'https://localhost:8080/countries';
  private devicesUrl = 'https://localhost:8080/devices';
  private testersUrl = 'https://localhost:8080/tester_match';

  constructor(private http: Http) { }

  public fetchCountries() {
    return this.http.get(this.countriesUrl)
      .map((res: Response) => res.json());
  }

  public fetchDevices() {
    return this.http.get(this.devicesUrl)
      .map((res: Response) => res.json());
  }

  public matchTesters(countries: string[], devices: string[]) {
    let params: URLSearchParams = new URLSearchParams();
    for (let country of countries) {
      params.append('country', country);
    }
    for (let device of devices) {
      params.append('device', device);
    }
    return this.http.get(this.testersUrl, {search: params})
      .map((res: Response) => res.json());
  }
}
