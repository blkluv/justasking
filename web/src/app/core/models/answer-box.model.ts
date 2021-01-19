import { BaseBoxModel } from './base-box.model'
import { QuestionModel } from './question.model'

export class AnswerBoxModel {
    Header?: string;
    Action?: string;
    Theme?: string;
    Questions?: QuestionModel[];
    IsPreview?: boolean;
    BaseBox?: BaseBoxModel;
}
