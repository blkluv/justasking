import { Injectable } from '@angular/core';
 
@Injectable()
export class ImageService {

  constructor() { }

  setDefaultProfileAvatar(event){ 
    event.target.src = `/assets/avatars/default.png`;
  } 

}