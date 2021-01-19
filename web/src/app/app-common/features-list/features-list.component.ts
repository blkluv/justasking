import { Component, Input, OnInit } from '@angular/core';

import { FeatureService } from '../../core/services/feature.service';

import { FeatureSectionModel } from '../../core/models/feature-section.model';

@Component({
  selector: 'app-features-list',
  templateUrl: './features-list.component.html',
  styleUrls: ['./features-list.component.scss']
})
export class FeaturesListComponent implements OnInit {

  @Input("featureSections") featureSections: FeatureSectionModel[];

  constructor(
    private featureService: FeatureService
  ) { }

  ngOnInit() {
    if (!this.featureSections) {
      this.featureSections = this.featureService.getBasicFeatureSections();
    }
  }
}
