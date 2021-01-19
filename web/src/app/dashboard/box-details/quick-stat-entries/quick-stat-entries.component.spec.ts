import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickStatEntriesComponent } from './quick-stat-entries.component';

describe('QuickStatEntriesComponent', () => {
  let component: QuickStatEntriesComponent;
  let fixture: ComponentFixture<QuickStatEntriesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickStatEntriesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickStatEntriesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
