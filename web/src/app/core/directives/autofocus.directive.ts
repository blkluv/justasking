import { Directive, OnInit, Input, ElementRef } from '@angular/core';

@Directive({
  selector: '[appAutofocus]'
})
export class AutofocusDirective {

  @Input('appAutofocus') enabled : boolean;

  constructor(private elementRef: ElementRef) { };

  ngOnInit(): void {
    if(this.enabled){
      this.elementRef.nativeElement.focus();
    }
  }

}
