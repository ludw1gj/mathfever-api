import React from 'react';

import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';

/** Terms page. */
const TermsView = () => {
  return (
    <div>
      <DocumentTitle>Terms of Use - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <PageHeading>Terms of Use</PageHeading>

      <ContentCard>
        <ul>
          <li>
            The information and the tools on this website are given as is
            without any warranty of any kind.
          </li>
          <li>
            This website, and its contributors are not be liable for any damages
            or losses resulting from the use of the content or the tools on this
            website.
          </li>
          <li>Using this website means you accept the Terms of Use.</li>
          <li>These Terms of Use are subject to change at any time.</li>
        </ul>
      </ContentCard>
    </div>
  );
};

export default TermsView;
