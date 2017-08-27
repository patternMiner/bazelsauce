import { Component } from '@angular/core';
import { DataService } from "./data.service";
import { OnInit } from '@angular/core';
import { SelectionModel, SelectionItem } from './selection.component';
import { TesterList, Tester } from './testers.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'TesterMatch';
  countries: SelectionItem[] = [];
  devices: SelectionItem[] = [];
  testers: Tester[] = [];

  testerList = new TesterList(this.testers);
  countrySelectionModel = new SelectionModel(this.countries);
  deviceSelectionModel = new SelectionModel(this.devices);

  constructor(private dataService: DataService){}

  ngOnInit() {
    this.dataService.fetchCountries().subscribe(data => {
      for (let item of data.Items) {
        this.countries.push(new SelectionItem(item, item, false));
      }
      console.info(data);
    });
    this.dataService.fetchDevices().subscribe(data => {
      for (let item of data.Items) {
        this.devices.push(new SelectionItem(item.Id, item.Description, false));
      }
      console.info(data);
    });
   }

  onSubmit() {
    let selectedCountries = this.getSelectedIds(this.countries);
    let selectedDevices = this.getSelectedIds(this.devices);
    console.info("Selected countries: ", selectedCountries);
    console.info("Selected devices: ", selectedDevices);
    this.testers.length = 0;
    this.dataService.matchTesters(selectedCountries, selectedDevices).subscribe(data => {
      this.testers.push(new Tester("Id", "FirstName", "LastName", "Country", "Rank"));
      for (let item of data.Items) {
        this.testers.push(new Tester(item.Id, item.FirstName, item.LastName, item.Country, item.Rank));
      }
    });
  }

  getSelectedIds(items: SelectionItem[]) {
    let selectedItems = [];
    for (let item of items) {
      if (item.selected) {
        selectedItems.push(item.id);
      }
    }
    return selectedItems;
  }
}
