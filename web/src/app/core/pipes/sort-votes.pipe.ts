import { Pipe, PipeTransform } from '@angular/core';
import { VotesBoxQuestionAnswerModel } from '../models/votes-box-question-answer.model';

@Pipe({
  name: 'sortVotes'
})
export class SortVotesPipe implements PipeTransform {

  transform(votes: VotesBoxQuestionAnswerModel[], sortFormat?: string):VotesBoxQuestionAnswerModel[] {
    let sortedVotes = votes;
    if(votes != null && typeof(votes) != 'undefined'){
      switch(sortFormat){
        case 'countAsc':
          sortedVotes = this.sortByCountAsc(votes);
          break;
        case 'countDesc':
          sortedVotes = this.sortByCountDesc(votes);
          break;
        case 'answerAsc':
          sortedVotes = this.sortByAnswerAsc(votes);
          break;
        case 'answerDesc':
          sortedVotes = this.sortByAnswerDesc(votes);
          break;
        default:
          break;
      }
    }
    return sortedVotes;
  }
  
  private sortByCountAsc(votes: VotesBoxQuestionAnswerModel[]):VotesBoxQuestionAnswerModel[]{
    let sortedVotes = votes.sort((a,b)=>{
      let responseA = a.Votes;
      let responseB = b.Votes;
      if(responseA<responseB)return -1;
      if(responseA>responseB)return 1;
      return 0;
    });
    return sortedVotes;
  }
  
  private sortByCountDesc(votes: VotesBoxQuestionAnswerModel[]):VotesBoxQuestionAnswerModel[]{
    let sortedVotes = votes.sort((a,b)=>{
      let responseA = a.Votes;
      let responseB = b.Votes;
      if(responseA<responseB)return 1;
      if(responseA>responseB)return -1;
      return 0;
    });
    return sortedVotes;
  }
  
  private sortByAnswerAsc(votes: VotesBoxQuestionAnswerModel[]):VotesBoxQuestionAnswerModel[]{
    let sortedVotes = votes.sort((a,b)=>{
      let responseA = a.Answer;
      let responseB = b.Answer;
      if(responseA<responseB)return -1;
      if(responseA>responseB)return 1;
      return 0;
    });
    return sortedVotes;
  }
  
  private sortByAnswerDesc(votes: VotesBoxQuestionAnswerModel[]):VotesBoxQuestionAnswerModel[]{
    let sortedVotes = votes.sort((a,b)=>{
      let responseA = a.Answer;
      let responseB = b.Answer;
      if(responseA<responseB)return 1;
      if(responseA>responseB)return -1;
      return 0;
    });
    return sortedVotes;
  }

}
