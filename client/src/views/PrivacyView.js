import React from 'react';

import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';
import SmallCaps from '../components/SmallCaps';

/** Privacy page. */
const PrivacyView = () => {
  return (
    <div>
      <DocumentTitle>Privacy - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <PageHeading>Privacy</PageHeading>

      <ContentCard>
        <h4>
          <strong>We use Google Analytics</strong>
        </h4>
        <p>
          Like many other websites, <SmallCaps>MathFever</SmallCaps> uses Google
          Analytics to gather statistics on the site's traffic in order to
          improve the quality of the website.
        </p>
      </ContentCard>
    </div>
  );
};

export default PrivacyView;
