export class QuestionBoxEntryModel {
    BoxId?: string;
    EntryId? : string;
    Code?: string;
    Question?: string;
    Upvotes?: number; 
    Downvotes?: number;
    IsHidden?: boolean;
    UserUpvoted?: boolean;
    UserDownvoted?: boolean;
    CreatedAt? : Date;
    CreatedBy?: number;
} 