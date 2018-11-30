import React, { Component } from 'react';
import { connect } from 'react-redux';
import styled from 'styled-components';

import ScrollLink from '../components/ScrollLink';
import LoadingSpinner from '../components/LoadingSpinner';
import ContentCard from '../components/ContentCard';
import mdlUpgrade from '../components/mdlUpgrade';
import { getCalculation, processCalculation } from '../util/query';
import { createSerializedAndValidatedForm } from '../util/form';
import ErrorView from './ErrorView';
import DocumentTitle from '../components/DocumentTitle';
import MetaDescription from '../components/MetaDescription';

const UserInputField = styled.div`
  margin-left: 5%;
`;

const CalculationButton = styled.button`
  margin-left: 5%;
  margin-top: 1%;
`;

const Heading = styled.h2`
  margin-left: 5%;
  margin-right: 5%;
  font-family: 'Roboto Thin', sans-serif;
  font-weight: 200;
`;

const CategoryHeadingLink = styled(ScrollLink)`
  margin-left: 5%;
  margin-right: 5%;
  font-family: 'Roboto Thin', sans-serif;
  font-weight: 200;
  text-decoration: none;
  color: inherit;
  &:hover {
    color: rgb(255, 64, 129);
  }
`;

// Styled used by server API
const StyledContentCard = styled(ContentCard)`
  .word-break {
    word-break: break-all;
  }

  .column-wrap {
    white-space: normal;
  }
`;

/** Calculation page. */
class CalculationView extends Component {
  state = {
    calculation: null,
    answer: null,
    error: null,
    loading: false
  };

  componentDidUpdate() {
    window.componentHandler.upgradeDom();
  }

  componentDidMount() {
    const { categorySlug } = this.props.match.params;
    const { calculationSlug } = this.props.match.params;
    const { categories } = this.props;

    getCalculation(categorySlug, calculationSlug, categories.data)
      .then(calculation => this.setState({ calculation }))
      .catch(error => {
        this.setState({ error: error.message });
      });
  }

  inputForm = calculation => {
    const submitInput = (e, calculationSlug) => {
      e.preventDefault();
      this.setState({ loading: true });

      const onError = error => {
        this.setState({ error, loading: false });
      };

      const form = e.target;
      const serializedFrom = createSerializedAndValidatedForm(form, onError);
      if (serializedFrom !== null) {
        // serialized form is valid and will be posted to the server.
        processCalculation(calculationSlug, serializedFrom)
          .then(answer => {
            this.setState({ answer, loading: false, error: null });
          })
          .catch(error => {
            this.setState({ error: error.message, loading: false });
          });
      }
    };

    return (
      <form
        id="calculation-form"
        onSubmit={e => submitInput(e, calculation.slug)}
      >
        {/* Input Field */}
        {calculation.inputInfo.map(inputField => {
          return (
            <UserInputField
              key={inputField.tag}
              className="mdl-textfield mdl-js-textfield mdl-textfield--floating-label"
            >
              <input
                className="mdl-textfield__input"
                name={inputField.tag}
                type="text"
                id={inputField.tag}
                autoComplete="off"
              />
              <label
                className="mdl-textfield__label"
                htmlFor={inputField.tag}
                id={`${inputField.tag}Label`}
              >
                {inputField.name}
              </label>
            </UserInputField>
          );
        })}

        {/* Button */}
        <CalculationButton
          type="submit"
          className="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect"
        >
          Calculate
        </CalculationButton>
      </form>
    );
  };

  render() {
    const { calculation, answer, loading, error } = this.state;

    if (!calculation && error) {
      return <ErrorView error={error} />;
    }
    if (!calculation) {
      return <LoadingSpinner />;
    }

    const content = () => {
      if (loading) {
        return <LoadingSpinner />;
      }
      if (error) return <ContentCard>{error}</ContentCard>;
      if (answer)
        return (
          <StyledContentCard>
            <div dangerouslySetInnerHTML={{ __html: answer }} />
          </StyledContentCard>
        );
    };

    return (
      <div>
        <DocumentTitle>{`${calculation.name} - MathFever`}</DocumentTitle>
        <MetaDescription>{calculation.description}</MetaDescription>

        {/* Heading */}
        <div className="calculation-page-heading">
          <h3>
            <CategoryHeadingLink to={calculation.categoryURL}>
              {calculation.category}:
            </CategoryHeadingLink>
          </h3>
          <Heading>{calculation.name}</Heading>
        </div>

        {/* Form */}
        {this.inputForm(calculation)}
        <hr className="separator" />

        {/* Content */}
        {content()}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    categories: state.categories
  };
};

export default connect(mapStateToProps)(mdlUpgrade(CalculationView));
