import React from 'react';

import DocumentTitle from '../components/DocumentTitle';
import ContentCard from '../components/ContentCard';
import goat from '../images/404-goat.jpg';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';
import ErrorImage from '../components/ErrorImage';
import MetaNoIndex from '../components/MetaNoIndex';

/** Not Found page. */
const NotFoundView = () => {
  return (
    <div>
      <DocumentTitle>Page Not Found - MathFever</DocumentTitle>
      <GenericDocMetaDesc />
      <MetaNoIndex />

      <ContentCard>
        <h2>Error 404 &rarr; This page doesn't exist.</h2>
        <h6>There's nothing here but us.</h6>
        <ErrorImage src={goat} alt="Error 404" />
      </ContentCard>
    </div>
  );
};

export default NotFoundView;
