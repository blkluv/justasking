export class QuestionBoxEntryVoteModel {
    Code?: string;
    EntryId?: string;
    VoteType?: string; 
    VoteValue?: number;
    IsUndo?: boolean;
    IsUpvoteFromDownvote?: boolean;
    IsDownvoteFromUpvote?: boolean;
    CreatedBy?: number;
}