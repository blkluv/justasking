import { AnswerBoxEntryModel } from './answer-box-entry.model'

export class QuestionModel {
    QuestionId: string;
    SortOrder: number;
    Question?: string;
    Entry?: string;
    Entries?: AnswerBoxEntryModel[];
    Collapsed?: boolean;
    IsActive?: boolean;
}
