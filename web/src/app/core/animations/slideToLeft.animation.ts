import {trigger, state, animate, style, transition} from '@angular/core';

export function slideToLeft() {
  return trigger('slideToLeft', [
    state('void', style({position:'inherit'}) ),
    state('*', style({position:'inherit'}) ),
    transition(':enter', [  // before 2.1: transition('void => *', [
      style({transform: 'translateX(100%)'}),
      animate('0.5s ease-in-out', style({transform: 'translateX(0%)'}))
    ]),
    transition(':leave', [  // before 2.1: transition('* => void', [
      style({transform: 'translateX(0%)'}),
      animate('0.5s ease-in-out', style({transform: 'translateX(-100%)'}))
    ])
  ]);
}