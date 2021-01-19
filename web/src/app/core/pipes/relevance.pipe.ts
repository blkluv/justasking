import { Pipe, PipeTransform } from '@angular/core';

import { QuestionBoxEntryModel } from '../models/question-box-entry.model';

@Pipe({
  name: 'relevance'
})
export class RelevancePipe implements PipeTransform {

  transform(items: QuestionBoxEntryModel[], args?: any): any {
    let sortedItems = items.sort((a: QuestionBoxEntryModel, b: QuestionBoxEntryModel) => {
        let aRelevance = a.Upvotes - a.Downvotes;
        let bRelevance = b.Upvotes - b.Downvotes;
        let aDate = a.CreatedAt;
        let bDate = b.CreatedAt;
        if (aRelevance > bRelevance) {
          return -1;
        } else if (aRelevance < bRelevance) {
          return 1;
        } else {
          if (aDate > bDate){
            return -1;
          } else if (aDate < bDate){
            return 1;
          }
          return 0;
        }
      });
      return sortedItems;
  }

}
