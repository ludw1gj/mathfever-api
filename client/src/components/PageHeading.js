import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

const Title = styled.h1`
  margin-left: 5%;
  margin-right: 5%;
  font-family: 'Roboto Thin', sans-serif;
  font-weight: 200;
`;

const Separator = styled.hr`
  margin-left: 1%;
  margin-right: 1%;
  border-top: 1px solid #f8f8f8;
  border-bottom: 1px solid rgba(0, 0, 0, 0.2);
`;

/** PageHeading component is a page heading with a separator. */
const PageHeading = ({ children }) => {
  return (
    <div>
      <Title>{children}</Title>
      <Separator />
    </div>
  );
};

PageHeading.propTypes = {
  children: PropTypes.string.isRequired
};

export default PageHeading;
