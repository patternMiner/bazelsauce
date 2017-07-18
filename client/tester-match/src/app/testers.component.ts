import { Component, Input } from '@angular/core';

@Component({
  selector: 'matched-testers',
  templateUrl: './testers.component.html',
  styleUrls: ['./testers.component.css']
})
export class MatchedTestersComponent {
  @Input() testerList : TesterList;
  constructor(){}
}

export class TesterList {
  constructor(public items: Tester[]){}
}

export class Tester {
  constructor(public id: Object, public firstName: Object, public lastName: Object,
      public country: Object, public rank: Object){}
}
