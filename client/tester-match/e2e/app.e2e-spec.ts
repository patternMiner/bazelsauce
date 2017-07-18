import { TesterMatchPage } from './app.po';

describe('tester-match App', () => {
  let page: TesterMatchPage;

  beforeEach(() => {
    page = new TesterMatchPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to TesterMatch!');
  });
});
