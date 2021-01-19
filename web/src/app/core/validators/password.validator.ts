import {AbstractControl} from '@angular/forms';
export class PasswordValidator {

    static MatchPassword(AC: AbstractControl) {
       let Password = AC.get('Password').value; // to get value in input tag
       let ConfirmPassword = AC.get('ConfirmPassword').value; // to get value in input tag
        if(Password != ConfirmPassword) {
            AC.get('ConfirmPassword').setErrors( {MatchPassword: true} )
        } else {
            return null
        }
    }
}