import React from 'react';

import ScrollLink from '../components/ScrollLink';
import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';

/** Help page. */
const HelpView = () => {
  return (
    <div>
      <DocumentTitle>Help - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <PageHeading>Help</PageHeading>

      <ContentCard>
        <h4>
          <strong>Found any mistakes?</strong>
        </h4>
        <p>
          If you find any mistakes post a comment on the{' '}
          <ScrollLink className={'no-underline-a-tag'} to={'/message-board'}>
            message board.
          </ScrollLink>{' '}
          detailing the issue.
        </p>

        <h4>
          <strong>Want to request a feature?</strong>
        </h4>
        <p>
          Post it on the{' '}
          <ScrollLink className={'no-underline-a-tag'} to={'/message-board'}>
            message board.
          </ScrollLink>{' '}
          _へ__(‾◡◝ )>
        </p>

        <h4>
          <strong>Do you want to ask a question?</strong>
        </h4>
        <p>
          Just ask on the{' '}
          <ScrollLink className={'no-underline-a-tag'} to={'/message-board'}>
            message board.
          </ScrollLink>
        </p>

        <h4>
          <strong>Got mad advice for the admin?</strong>
        </h4>
        <p>
          →{' '}
          <ScrollLink className={'no-underline-a-tag'} to={'/message-board'}>
            message board.
          </ScrollLink>{' '}
          ←
        </p>
      </ContentCard>
    </div>
  );
};

export default HelpView;
