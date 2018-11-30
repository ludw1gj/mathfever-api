import React, { Component } from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';

import networkingImage from '../images/category/networking.jpg';
import primesAndFactorsImage from '../images/category/primes-and-factors.jpg';
import percentagesImage from '../images/category/percentages.jpg';
import totalSurfaceAreaImage from '../images/category/total-surface-area.jpg';

import ScrollLink from '../components/ScrollLink';
import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import ContentCard from '../components/ContentCard';
import LoadingSpinner from '../components/LoadingSpinner';
import mdlUpgrade from '../components/mdlUpgrade';
import { getCategory } from '../util/query';
import ErrorView from '../views/ErrorView';
import MetaDescription from '../components/MetaDescription';

/** Category page. */
class CategoryView extends Component {
  state = {
    category: null,
    error: null
  };

  images = new Map([
    ['networking', networkingImage],
    ['primes-and-factors', primesAndFactorsImage],
    ['percentages', percentagesImage],
    ['total-surface-area', totalSurfaceAreaImage]
  ]);

  componentDidUpdate() {
    window.componentHandler.upgradeDom();
  }

  componentDidMount() {
    const { categorySlug } = this.props.match.params;
    const { categories } = this.props;

    getCategory(categorySlug, categories.data)
      .then(category => this.setState({ category }))
      .catch(err => {
        this.setState({ error: err.message });
      });
  }

  render() {
    const { category, error } = this.state;

    if (!category && error) {
      return <ErrorView error={error} />;
    }

    const StyleScrollLink = styled(ScrollLink)`
      text-decoration: none;
      &:hover {
        text-decoration: underline;
      }
    `;

    const categoryContent = () => {
      return (
        <div>
          <PageHeading>{category.name}</PageHeading>
          <MetaDescription>{category.description}</MetaDescription>

          <ContentCard>
            <h4>{category.name} Calculations:</h4>
            <ul>
              {category.calculations.map(calc => {
                return (
                  <li key={calc.slug}>
                    <StyleScrollLink to={calc.url}>{calc.name}</StyleScrollLink>
                  </li>
                );
              })}
            </ul>

            {category.calculations.map(calc => {
              return (
                <div key={calc.slug}>
                  <hr />
                  <h4>{calc.name}</h4>
                  <div dangerouslySetInnerHTML={{ __html: calc.example }} />
                </div>
              );
            })}
          </ContentCard>
        </div>
      );
    };

    return (
      <div>
        <DocumentTitle>
          {category ? `${category.name} - MathFever` : 'Category - MathFever'}
        </DocumentTitle>

        {category ? categoryContent() : <LoadingSpinner />}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    categories: state.categories
  };
};

export default connect(mapStateToProps)(mdlUpgrade(CategoryView));
