import React from 'react';
import PropTypes from 'prop-types';

import DocumentTitle from '../components/DocumentTitle';
import ContentCard from '../components/ContentCard';
import panda from '../images/500-panda.jpg';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';
import ErrorImage from '../components/ErrorImage';

/** Error page. */
const ErrorView = ({ error }) => {
  return (
    <div>
      <DocumentTitle>Error - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <ContentCard>
        <h2>Error &rarr; Something went wrong.</h2>
        <p>
          <b>{error}</b>
        </p>
        <h6>Here's a panda:</h6>
        <ErrorImage src={panda} alt="Error 500" />
      </ContentCard>
    </div>
  );
};

ErrorView.propTypes = {
  error: PropTypes.string.isRequired
};

export default ErrorView;
