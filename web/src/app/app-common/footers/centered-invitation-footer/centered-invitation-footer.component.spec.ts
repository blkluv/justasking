import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CenteredInvitationFooterComponent } from './centered-invitation-footer.component';

describe('CenteredInvitationFooterComponent', () => {
  let component: CenteredInvitationFooterComponent;
  let fixture: ComponentFixture<CenteredInvitationFooterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CenteredInvitationFooterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CenteredInvitationFooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
