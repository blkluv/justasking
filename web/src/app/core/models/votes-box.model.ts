import { BaseBoxModel } from './base-box.model'
import { VotesBoxQuestionModel } from './votes-box-question.model';
import { VotesBoxQuestionAnswerModel } from './votes-box-question-answer.model';

export class VotesBoxModel {
    BoxId?: string;
    BaseBox?: BaseBoxModel;
    Questions?: VotesBoxQuestionModel[];
    IsPreview?: boolean;
}
