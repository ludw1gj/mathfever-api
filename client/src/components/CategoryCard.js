import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

import ScrollLink from '../components/ScrollLink';

const StyleScrollLink = styled(ScrollLink)`
  text-decoration: none;
  color: inherit;
  &:hover {
    text-decoration: underline;
  }
`;

const Card = styled.div`
  border: solid 1px #e4e4e4;
`;

/** CategoryCard component is a card containing information about a category. */
const CategoryCard = ({ image, category }) => {
  const CardTitle = styled.div`
    background: linear-gradient(rgba(0, 0, 0, 0.3), rgba(0, 0, 0, 0.3)),
      url(${image}) center/cover;
    color: #fff;
    height: 176px;
  `;

  const cardText = category => {
    return category.calculations.map(calc => {
      return (
        <li key={calc.slug}>
          <StyleScrollLink to={calc.url}>{calc.name}</StyleScrollLink>
        </li>
      );
    });
  };

  return (
    <Card className="mdl-card mdl-cell mdl-cell--top mdl-cell--4 mdl-cell--4-col-phone">
      <CardTitle className="mdl-card__title">
        <h6 className="mdl-card__title-text">
          <StyleScrollLink to={category.url}>{category.name}</StyleScrollLink>
        </h6>
      </CardTitle>
      <div className="mdl-card__supporting-text">
        <ul>{cardText(category)}</ul>
      </div>
    </Card>
  );
};

CategoryCard.propTypes = {
  image: PropTypes.string.isRequired,
  category: PropTypes.object.isRequired
};

export default CategoryCard;
