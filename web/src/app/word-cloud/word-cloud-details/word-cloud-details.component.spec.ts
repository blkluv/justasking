import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordCloudDetailsComponent } from './word-cloud-details.component';

describe('WordCloudDetailsComponent', () => {
  let component: WordCloudDetailsComponent;
  let fixture: ComponentFixture<WordCloudDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordCloudDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordCloudDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
