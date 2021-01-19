import { Component, Input, OnInit, OnChanges, SimpleChanges, HostListener } from '@angular/core';

declare var WordCloud:any;

@Component({
  selector: 'app-word-cloud',
  templateUrl: './word-cloud.component.html',
  styleUrls: ['./word-cloud.component.scss']
})
export class WordCloudComponent implements OnInit, OnChanges {

  @Input("defaultWord") defaultWord: string;
  @Input("data") data : [[string,number]];

  constructor() {
    this.data = this.data || [[this.defaultWord,110]];
  }

  @HostListener('window:resize', ['$event'])
  onResize(event) {
    let currentWidth = event.target.innerWidth;

    if(this.data && this.data.length>0){
      let newDefaultWordSize = this.getDefaultWordSizeByWindowWidth(currentWidth);
      this.data[0][1] = newDefaultWordSize;
    }

    this.createWordCloud(this.data);
  }

  getDefaultWordSizeByWindowWidth(width:number):number{
    let size = 110;
    
    if(width < 768){
      size = 40;
    }else if(width < 993){
      size = 60;
    }else if(width < 1185){
      size = 80;
    } 

    return size;
  }
  
  private createWordCloud(data : [[string,number]]){
    let options = { 
      list: data,
      shape: "circle",
      shuffle: false,
      backgroundColor: "transparent",
      background: "transparent",
      color: "white",
      rotateRatio: 0,
      minRotation: 1,
      maxRotation: 1, 
    }
    let el = document.getElementById('wordcloud');
    WordCloud(el, options);
  }

  ngOnChanges(changes: SimpleChanges) {
    if(changes["defaultWord"] && !changes["defaultWord"].firstChange){
      let newDefaultWord = changes["defaultWord"].currentValue;
      let defaultWordSize = this.getDefaultWordSizeByWindowWidth(window.innerWidth);
      this.data[0] = [newDefaultWord,defaultWordSize];
      this.createWordCloud(this.data);
    }else if(changes["data"] && !changes["data"].firstChange){
      this.createWordCloud(this.data);
    }
  } 

  ngOnInit() {
    this.createWordCloud(this.data);
  }

}
