import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordCloudEntryComponent } from './word-cloud-entry.component';

describe('WordCloudEntryComponent', () => {
  let component: WordCloudEntryComponent;
  let fixture: ComponentFixture<WordCloudEntryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordCloudEntryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordCloudEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
