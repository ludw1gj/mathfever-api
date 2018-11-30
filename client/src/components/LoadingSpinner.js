import React from 'react';
import styled from 'styled-components';

const Container = styled.div`
  padding-top: 3em;
  margin: auto;
  text-align: center;
`;

/** LoadingSpinner component is an MDL spinner. */
const LoadingSpinner = () => {
  return (
    <Container>
      <div className="mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active" />
    </Container>
  );
};

export default LoadingSpinner;
