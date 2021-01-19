import { Injectable } from '@angular/core';

@Injectable()
export class ProcessingService {

  processingAnimationIsActive: boolean;

  constructor() { 
    this.processingAnimationIsActive = false;
  }

  toggleProcessingAnimation(){
    this.processingAnimationIsActive = !this.processingAnimationIsActive;
  }

  enableProcessingAnimation(){
    this.processingAnimationIsActive = true;
  }

  disableProcessingAnimation(){
    this.processingAnimationIsActive = false;
  }

}
