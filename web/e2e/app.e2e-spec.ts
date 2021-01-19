import { JustAskingPage } from './app.po';

describe('just-asking App', () => {
  let page: JustAskingPage;

  beforeEach(() => {
    page = new JustAskingPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
