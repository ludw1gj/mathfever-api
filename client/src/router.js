import React from 'react';
import { Switch, Route } from 'react-router-dom';

import AboutView from './views/AboutView';
import HomeView from './views/HomeView';
import HelpView from './views/HelpView';
import PrivacyView from './views/PrivacyView';
import TermsView from './views/TermsView';
import NotFoundView from './views/NotFoundView';
import ConversionTableView from './views/ConversionTableView';
import CategoryView from './views/CategoryView';
import MessageBoardView from './views/MessageBoardView';
import CalculationView from './views/CalculationView';

/** AppRouter component handles routing. */
const AppRouter = () => {
  return (
    <Switch>
      <Route exact path="/" component={HomeView} />
      <Route exact path="/about" component={AboutView} />
      <Route exact path="/help" component={HelpView} />
      <Route exact path="/privacy" component={PrivacyView} />
      <Route exact path="/terms" component={TermsView} />
      <Route exact path="/message-board" component={MessageBoardView} />
      <Route
        exact
        path="/category/networking/conversion-table"
        component={ConversionTableView}
      />
      <Route exact path="/category/:categorySlug" component={CategoryView} />
      <Route
        exact
        path="/category/:categorySlug/:calculationSlug"
        component={CalculationView}
      />

      <Route path="*" component={NotFoundView} />
    </Switch>
  );
};

export default AppRouter;
