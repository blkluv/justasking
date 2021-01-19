import { VotesBoxQuestionAnswerModel } from './votes-box-question-answer.model'

export class VotesBoxQuestionModel {
    Answers?: VotesBoxQuestionAnswerModel[];
    BoxId?: string;
    Header?: string;
    IsActive?: boolean; 
    QuestionId?: string;
    SortOrder?: number; 
    SelectedAnswerId?: string;
}
