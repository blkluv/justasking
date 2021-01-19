import { Component, Input, OnInit, OnChanges, SimpleChanges } from '@angular/core';

import { VotesBoxQuestionAnswerModel } from '../core/models/votes-box-question-answer.model';

@Component({
  selector: 'app-bar-graph',
  templateUrl: './bar-graph.component.html',
  styleUrls: ['./bar-graph.component.scss']
})
export class BarGraphComponent implements OnInit {

  @Input("theme") theme: string;
  @Input("header") header: string;
  @Input("showTotalCount") showTotalCount: boolean;
  @Input("answers") answers: VotesBoxQuestionAnswerModel[];
  totalVotesCount: number;

  constructor() {
    this.totalVotesCount = 0;
    this.header = null;
  }

  ngOnChanges(changes: SimpleChanges) {
    if(changes["answers"] && !changes["answers"].firstChange){
      this.totalVotesCount = this.calculateTotalVotesCount(changes["answers"].currentValue);
    }
  }

  private calculateTotalVotesCount(answers: VotesBoxQuestionAnswerModel[]): number {
    let totalVotesCount = 0;
    answers.forEach(answer => {
      totalVotesCount += answer.Votes;
    });
    return totalVotesCount;
  }

  ngOnInit() {
    this.totalVotesCount = this.calculateTotalVotesCount(this.answers);
  }
}
