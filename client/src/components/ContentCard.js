import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

const Content = styled.div`
  margin: auto;
  padding: 20px 5%;
  overflow: auto;
  box-shadow: 0 3px 4px rgba(0, 0, 0, 0.15);
`;

/** ContentCard component is a card which childs content. */
const ContentCard = ({ className, children }) => {
  return (
    <div className={`mdl-grid ${className}`}>
      <Content className="mdl-color--white content mdl-color-text--grey-800 mdl-cell mdl-cell--9-col">
        {children}
      </Content>
    </div>
  );
};

ContentCard.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
};

export default ContentCard;
