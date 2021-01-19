import { QuestionModel } from './question.model';

export class AnswerBoxPresentationModel {
    BoxId?: string;
    Code?: string;
    Questions?: QuestionModel[];
    IsHidden?: boolean;
    CreatedAt? : Date;
    CreatedBy?: number;
}
    