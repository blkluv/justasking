import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HttpClientJsonpModule } from '@angular/common/http';

import { MatMenuModule } from '@angular/material';
import { MatIconModule } from '@angular/material';
import { MatDialogModule } from '@angular/material';
import { MatCardModule } from '@angular/material';
import { MatButtonModule } from '@angular/material';
import { MatOptionModule } from '@angular/material';
import { MatDatepickerModule } from '@angular/material';
import { MatToolbarModule } from '@angular/material';
import { MatCheckboxModule } from '@angular/material';
import { MatSidenavModule } from '@angular/material';
import { MatSnackBarModule } from '@angular/material';
import { MatInputModule } from '@angular/material';
import { MatSelectModule } from '@angular/material';
import { MatFormFieldModule } from '@angular/material';
import { MatTooltipModule } from '@angular/material';
import { MatTabsModule } from '@angular/material';
import { MatListModule } from '@angular/material';
import { MatNativeDateModule } from '@angular/material';
import { MatProgressSpinnerModule } from '@angular/material';
import { MatProgressBarModule } from '@angular/material';
import { MatSlideToggleModule } from '@angular/material';
import { MatButtonToggleModule } from '@angular/material';

import { ClipboardModule } from 'ngx-clipboard';
import { ShareModule } from '@ngx-share/core';

import { AutofocusDirective } from '../core/directives/autofocus.directive';
import { AutosizeDirective } from '../core/directives/autosize.directive';
import { AlphaNumericsOnlyDirective } from '../core/directives/alpha-numerics-only.directive';

import { ThemeByIndexPipe } from '../core/pipes/theme-by-index.pipe';
import { RelevancePipe } from '../core/pipes/relevance.pipe';
import { UniqueEntriesPipe } from '../core/pipes/unique-entries.pipe';
import { PhoneNumberPipe } from '../core/pipes/phone-number.pipe';
import { SortVotesPipe } from '../core/pipes/sort-votes.pipe';

import { TermsAndConditionsComponent } from '../terms-and-conditions/terms-and-conditions.component';
import { PageNotFoundComponent } from '../page-not-found/page-not-found.component';
import { InternalServerErrorComponent } from '../internal-server-error/internal-server-error.component';
import { UserMenuComponent } from './user-menu/user-menu.component';
import { EntryTopnavComponent } from './topnavs/entry-topnav/entry-topnav.component';
import { PresentationTopnavComponent } from './topnavs/presentation-topnav/presentation-topnav.component';
import { CenteredInvitationFooterComponent } from './footers/centered-invitation-footer/centered-invitation-footer.component';
import { QuickStatBoxComponent } from '../dashboard/box-details/quick-stat-box/quick-stat-box.component';
import { QuickStatParticipantsComponent } from '../dashboard/box-details/quick-stat-participants/quick-stat-participants.component';
import { QuickStatEntriesComponent } from '../dashboard/box-details/quick-stat-entries/quick-stat-entries.component';
import { BarGraphComponent } from '../bar-graph/bar-graph.component';
import { AccountUpgradePromptDialogComponent } from './account-upgrade-prompt-dialog/account-upgrade-prompt-dialog.component';
import { PlanPricingCardComponent } from './plan-pricing-card/plan-pricing-card.component';
import { TermsAndConditionsDialogComponent } from './terms-and-conditions-dialog/terms-and-conditions-dialog.component'
import { ConfirmationDialogComponent } from './confirmation-dialog/confirmation-dialog.component'
import { SupportWidgetComponent } from './support-widget/support-widget.component';
import { NoAccessComponent } from './no-access/no-access.component';
import { FeaturesListComponent } from './features-list/features-list.component';

@NgModule({
  imports: [
    MatMenuModule,
    MatIconModule,
    MatDialogModule,
    MatCardModule,
    MatButtonModule,
    MatOptionModule,
    MatDatepickerModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatSidenavModule,
    MatSnackBarModule,
    MatInputModule,
    MatSelectModule,
    MatFormFieldModule,
    MatTooltipModule,
    MatTabsModule,
    MatListModule,
    MatNativeDateModule,
    MatProgressSpinnerModule,
    MatProgressBarModule,
    MatSlideToggleModule,
    MatButtonToggleModule,

    FormsModule,
    ReactiveFormsModule,
    CommonModule,
    ClipboardModule,
    RouterModule,

    HttpClientModule,       // for share counts
    HttpClientJsonpModule,  // for linkedin and tumblr share counts
    ShareModule.forRoot()
  ],
  exports: [
    MatNativeDateModule,
    MatProgressSpinnerModule,
    MatProgressBarModule,
    MatSlideToggleModule,
    MatButtonToggleModule,

    MatMenuModule,
    MatIconModule,
    MatDialogModule,
    MatCardModule,
    MatButtonModule,
    MatOptionModule,
    MatDatepickerModule,
    MatToolbarModule,
    MatCheckboxModule,
    MatSidenavModule,
    MatSnackBarModule,
    MatInputModule,
    MatSelectModule,
    MatFormFieldModule,
    MatTooltipModule,
    MatTabsModule,
    MatListModule,

    RouterModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    ClipboardModule,
    ThemeByIndexPipe,
    RelevancePipe,
    UniqueEntriesPipe,
    PhoneNumberPipe,
    SortVotesPipe,
    AutofocusDirective,
    AutosizeDirective,
    AlphaNumericsOnlyDirective,
    TermsAndConditionsComponent,
    PageNotFoundComponent,
    InternalServerErrorComponent,
    UserMenuComponent,
    EntryTopnavComponent,
    PresentationTopnavComponent,
    CenteredInvitationFooterComponent,
    QuickStatBoxComponent,
    QuickStatParticipantsComponent,
    QuickStatEntriesComponent,
    BarGraphComponent,
    AccountUpgradePromptDialogComponent,
    SupportWidgetComponent,
    NoAccessComponent,
    TermsAndConditionsDialogComponent,
    ConfirmationDialogComponent,
    FeaturesListComponent,
    PlanPricingCardComponent,
  ],
  declarations: [
    ThemeByIndexPipe,
    RelevancePipe,
    UniqueEntriesPipe,
    PhoneNumberPipe,
    SortVotesPipe,
    AutofocusDirective,
    AutosizeDirective,
    AlphaNumericsOnlyDirective,
    TermsAndConditionsComponent,
    PageNotFoundComponent,
    InternalServerErrorComponent,
    UserMenuComponent,
    EntryTopnavComponent,
    PresentationTopnavComponent,
    CenteredInvitationFooterComponent,
    QuickStatBoxComponent,
    QuickStatParticipantsComponent,
    QuickStatEntriesComponent,
    BarGraphComponent,
    AccountUpgradePromptDialogComponent,
    SupportWidgetComponent,
    NoAccessComponent,
    TermsAndConditionsDialogComponent,
    ConfirmationDialogComponent,
    FeaturesListComponent,
    PlanPricingCardComponent,
  ],
  entryComponents: [
    AccountUpgradePromptDialogComponent,
    TermsAndConditionsDialogComponent,
    ConfirmationDialogComponent,
  ]
})
export class AppCommonModule { }
