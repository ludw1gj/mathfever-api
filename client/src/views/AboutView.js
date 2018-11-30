import React from 'react';

import ScrollLink from '../components/ScrollLink';
import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';
import SmallCaps from '../components/SmallCaps';

/** About page. */
const AboutView = () => {
  return (
    <div>
      <DocumentTitle>About - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <PageHeading>About</PageHeading>

      <ContentCard>
        <h4>
          <strong>Why, hello there!</strong>
        </h4>
        <p>
          This website <SmallCaps>MathFever</SmallCaps> aims to help you do a
          dodgy job on your homework. A lone admin built this site so he could
          help people spend less time on their homework and more time out with
          friends. Users can find mathematical proof and answers to common math
          problems, with values of their choosing. Math problem such as
          Highest/Lowest Common Denominator, Volume/Total Surface Area of a
          handful of different shapes, Conversions for the field of Networking
          (Binary, Decimal, and Hexadecimal), and more. The admin hopes you find
          this site a somewhat bit useful.
        </p>
        <p>
          If you see any bugs, report them and they will be squashed! If you
          want to request any features or have ideas for improvement don't
          hesitate to give the admin a yell, he likes to program.
        </p>
        <h5>
          <strong>Contact</strong>
        </h5>
        <p>
          Post a comment on the{' '}
          <ScrollLink className={'no-underline-a-tag'} to={'/message-board'}>
            message board.
          </ScrollLink>
        </p>
      </ContentCard>
    </div>
  );
};

export default AboutView;
