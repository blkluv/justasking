import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WordCloudPresentationComponent } from './word-cloud-presentation.component';

describe('WordCloudPresentationComponent', () => {
  let component: WordCloudPresentationComponent;
  let fixture: ComponentFixture<WordCloudPresentationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WordCloudPresentationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WordCloudPresentationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
