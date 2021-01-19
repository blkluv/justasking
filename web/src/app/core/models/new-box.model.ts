import { BaseBoxModel } from './base-box.model'
import { VotesBoxQuestionAnswerModel } from './votes-box-question-answer.model';
import { QuestionModel } from './question.model';

export class NewBoxModel {
    Header?:string;
    DefaultWord?:string;
    Theme?:string; 
    Action?: string;
    Answers?: VotesBoxQuestionAnswerModel[];
    Question?: string;
    Questions?: QuestionModel[];
    IsPreview?: boolean; 
    BaseBox?: BaseBoxModel;
}
