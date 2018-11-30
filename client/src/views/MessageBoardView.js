import React from 'react';
import styled from 'styled-components';

import cat from '../images/first-post.jpg';

import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';

const RoundImage = styled.img`
  max-width: 100%;
  border-radius: 15px;
`;

/** Message Board page. */
const MessageBoardView = () => {
  return (
    <div>
      <DocumentTitle>Message Board - MathFever</DocumentTitle>
      <GenericDocMetaDesc />

      <PageHeading>Message Board</PageHeading>

      <ContentCard>
        <div className="post">
          <h3>A Warm Welcome To This Slightly Terrible Site!</h3>
          <p>
            May it be slightly terrible or just plain terrible, if at least
            someone finds this site useful I'll be as happy as a clam.
          </p>
          <p>What's a clam, here's a cat:</p>
          <RoundImage src={cat} width="400" alt="cat" />
          <p>
            <strong>Nov. 19, 2016</strong>
          </p>
          <p>Page 1 of 1</p>
        </div>
      </ContentCard>
    </div>
  );
};

export default MessageBoardView;
