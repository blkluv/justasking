import { Directive, ElementRef } from '@angular/core';

@Directive({
  selector: '[appAlphaNumericsOnly]'
})
export class AlphaNumericsOnlyDirective {

    constructor(public el: ElementRef) {
        this.el.nativeElement.onkeypress = (evt) => {
          var inp = String.fromCharCode(evt.keyCode);
          if (!(/[a-zA-Z0-9]/.test(inp))){
              //input was not a letter, number, hyphen, underscore or space");
              evt.preventDefault();
          }
        };
    }

}
