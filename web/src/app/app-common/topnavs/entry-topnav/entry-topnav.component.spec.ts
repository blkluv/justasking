import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EntryTopnavComponent } from './entry-topnav.component';

describe('EntryTopnavComponent', () => {
  let component: EntryTopnavComponent;
  let fixture: ComponentFixture<EntryTopnavComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EntryTopnavComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EntryTopnavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
