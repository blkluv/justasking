import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'uniqueEntries'
})
export class UniqueEntriesPipe implements PipeTransform {

  transform(entries: any[], args?: any): any {
    let uniqueEntries:any[] = [];
    let uniqueEntriesFlags = {};

    if(entries && typeof(entries) == 'object'){
      entries.forEach(function(entry, index){
        let response = entry.Response || entry.Question;
        if(!uniqueEntriesFlags[response]){
          uniqueEntriesFlags[response] = { Count : 1, Index: uniqueEntries.length};
          entry.Response = entry.Response || entry.Question;
          entry.Count = 1;
          uniqueEntries.push(entry);
        }else{
          let nonUniqueEntryIndex = uniqueEntriesFlags[response].Index;
          uniqueEntriesFlags[response].Count = uniqueEntriesFlags[response].Count + 1; 
          uniqueEntries[nonUniqueEntryIndex].Count = uniqueEntriesFlags[response].Count; 
        }
      })
    }

    // uniqueEntries = uniqueEntries.sort((a,b)=>{
    //   let responseA = a.Response;
    //   let responseB = b.Response;
    //   if(responseA < responseB) return -1;
    //   if(responseA > responseB)return 1;
    //   return 0;
    // });

    return uniqueEntries;
  }

}
