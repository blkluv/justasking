import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'phoneNumber'
})
export class PhoneNumberPipe implements PipeTransform {

  transform(text: string, format?: string): string {
    let phoneNumberText = text;
    if(text && text.length >= 10){
      let countryCode = text.substring(0, text.length - 10);
      let first3 = text.substring(text.length-10, text.length - 7);
      let middle3 = text.substring(text.length-7, text.length - 4);
      let last4 = text.substring(text.length - 4);
      if(format && format == 'country'){
        phoneNumberText = `${countryCode} (${first3}) ${middle3}-${last4}`;
      }else if(format && format == 'pound'){
        phoneNumberText = `# (${first3}) ${middle3}-${last4}`;
      }else if(format && format == 'default'){
        phoneNumberText = `(${first3}) ${middle3}-${last4}`;
      }else{
        console.error("PhoneNumberPipe was not given a format.");
      }
    }else{
      phoneNumberText="";
    }
    return phoneNumberText;
  }

}
