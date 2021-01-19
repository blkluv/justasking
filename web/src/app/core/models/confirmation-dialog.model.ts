import { Observable } from 'rxjs/Observable';

export class ConfirmationDialogModel {
    Header?:string;
    Texts?:string[];
    CancelText?:string;
    CancelAction?: ()=> Observable<any>;
    ConfirmText?:string;
    ConfirmAction?: ()=> Observable<any>;
}