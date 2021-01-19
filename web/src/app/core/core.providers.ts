import { AuthGuard } from './security/auth.guard'; 

import { UserService } from './services/user.service'; 
import { BaseApiService } from './services/base-api.service'; 
import { ColorService } from './services/color.service'; 
import { SidenavService } from './services/sidenav.service'; 
import { RandomizerService } from './services/randomizer.service'; 
import { AnswerBoxService } from './services/answer-box.service'; 
import { BoxService } from './services/box.service'; 
import { QuestionBoxService } from './services/question-box.service'; 
import { WordCloudService } from './services/word-cloud.service'; 
import { IdpAuthenticationService } from './services/idp-authentication.service'; 
import { ProcessingService } from './services/processing.service'; 
import { WebSocketService } from './services/web-socket.service'; 
import { GoogleAnalyticsService } from './services/google-analytics.service';
import { NotificationsService } from './services/notifications.service';
import { VotesBoxService } from './services/votes-box.service';
import { PlanService } from './services/plan.service';
import { AccountService } from './services/account.service';
import { StripeService } from './services/stripe.service';
import { SupportService } from './services/support.service';
import { ReleaseService } from './services/release.service';
import { FeatureRequestService } from './services/feature-request.service';
import { ImageService } from './services/image.service';
import { FeatureService } from './services/feature.service';

import { BaseRepoService } from './repos/base-repo.service';
import { QuestionBoxEntryVoteRepoService } from './repos/question-box-entry-vote-repo.service';
import { VotesBoxVoteRepoService } from './repos/votes-box-vote-repo.service';

import { RelevancePipe } from './pipes/relevance.pipe';
import { UniqueEntriesPipe } from './pipes/unique-entries.pipe';
import { SortVotesPipe } from './pipes/sort-votes.pipe';

export const CORE_PROVIDERS = [
    //Guards
    AuthGuard,

    //Services
    BaseApiService,
    UserService,
    ColorService,
    SidenavService,
    IdpAuthenticationService,
    RandomizerService,
    BoxService,
    QuestionBoxService,
    AnswerBoxService,
    WordCloudService,
    ProcessingService,
    WebSocketService,
    NotificationsService,
    GoogleAnalyticsService,
    VotesBoxService,
    PlanService,
    AccountService,
    StripeService,
    SupportService,
    ReleaseService,
    FeatureRequestService,
    ImageService,
    FeatureService,

    //Repo Services
    BaseRepoService,
    QuestionBoxEntryVoteRepoService,
    VotesBoxVoteRepoService,

    //Pipes
    RelevancePipe,
    UniqueEntriesPipe,
    SortVotesPipe
];