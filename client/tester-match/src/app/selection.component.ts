import { Component, Input } from '@angular/core';

@Component({
  selector: 'selection-list',
  templateUrl: './selection.component.html',
  styleUrls: ['./selection.component.css']
})
export class SelectionListComponent {
  @Input() selectionModel : SelectionModel;
  constructor(){}
}

export class SelectionModel {
  constructor(public items: SelectionItem[]){}

  selectAll() {
    for (let item of this.items) {
      item.selected = true;
    }
  }

  clearSelection() {
    for (let item of this.items) {
      item.selected = false;
    }
  }
}

export class SelectionItem {
  constructor(public id: Object, public label: Object, public selected: boolean){}
}
