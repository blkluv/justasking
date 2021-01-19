import { BaseBoxModel } from './base-box.model'

export class QuestionBoxModel {
    BoxId?: string;
    Header?: string;
    Action?: string;
    Theme?: string;
    Question?: string;
    IsPreview?: boolean;
    BaseBox?: BaseBoxModel;
}
