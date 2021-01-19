import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'themeByIndex'
})
export class ThemeByIndexPipe implements PipeTransform {

  transform(index: number, args?: any): any {
    let theme = "";

    if(index > 18){
      index = index%19;
    }
    switch (index) {
      case 0:{
        theme = "core-red";
        break;
      }
      case 1:{
        theme = "core-pink";
        break;
      }
      case 2:{
        theme = "core-purple";
        break;
      }
      case 3:{
        theme = "core-deep-purple";
        break;
      }
      case 4:{
        theme = "core-indigo";
        break;
      }
      case 5:{
        theme = "core-blue";
        break;
      }
      case 6:{
        theme = "core-light-blue";
        break;
      }
      case 7:{
        theme = "core-cyan";
        break;
      }
      case 8:{
        theme = "core-teal";
        break;
      }
      case 9:{
        theme = "core-green";
        break;
      }
      case 10:{
        theme = "core-light-green";
        break;
      }
      case 11:{
        theme = "core-lime";
        break;
      }
      case 12:{
        theme = "core-yellow";
        break;
      }
      case 13:{
        theme = "core-amber";
        break;
      }
      case 14:{
        theme = "core-orange";
        break;
      }
      case 15:{
        theme = "core-deep-orange";
        break;
      }
      case 16:{
        theme = "core-brown";
        break;
      }
      case 17:{
        theme = "core-grey";
        break;
      }
      case 18:{
        theme = "core-blue-grey";
        break;
      }
      default:
        break;
    }

    return theme;
  }
}
