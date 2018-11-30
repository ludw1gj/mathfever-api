import React, { Component } from 'react';
import { connect } from 'react-redux';
import { getCategories } from '../store/actions/categoriesActions';

import networkingImage from '../images/category/networking.jpg';
import primesAndFactorsImage from '../images/category/primes-and-factors.jpg';
import percentagesImage from '../images/category/percentages.jpg';
import totalSurfaceAreaImage from '../images/category/total-surface-area.jpg';

import DocumentTitle from '../components/DocumentTitle';
import PageHeading from '../components/PageHeading';
import CategoryCard from '../components/CategoryCard';
import LoadingSpinner from '../components/LoadingSpinner';
import mdlUpgrade from '../components/mdlUpgrade';
import GenericDocMetaDesc from '../components/GenericDocMetaDesc';

/** Help page. */
class HomeView extends Component {
  images = new Map([
    ['networking', networkingImage],
    ['primes-and-factors', primesAndFactorsImage],
    ['percentages', percentagesImage],
    ['total-surface-area', totalSurfaceAreaImage]
  ]);

  componentDidMount() {
    if (!this.props.categories.data) this.props.getCategories();
  }

  render() {
    const { categories } = this.props;

    const categoryCards = categories => {
      const cards = categories.map(category => {
        return (
          <CategoryCard
            key={category.slug}
            image={this.images.get(category.slug)}
            category={category}
          />
        );
      });
      return <div className="mdl-grid">{cards}</div>;
    };

    return (
      <div>
        <DocumentTitle>MathFever</DocumentTitle>
        <GenericDocMetaDesc />

        <PageHeading>Home</PageHeading>

        {categories.data && Array.isArray(categories.data) ? (
          categoryCards(categories.data)
        ) : (
          <LoadingSpinner />
        )}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    categories: state.categories
  };
};

const mapDispatchToProps = dispatch => {
  return {
    getCategories: () => dispatch(getCategories())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(mdlUpgrade(HomeView));
