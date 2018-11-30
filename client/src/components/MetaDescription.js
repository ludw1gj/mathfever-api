import React from 'react';
import PropTypes from 'prop-types';

/** MetaDescription component sets the document meta description. */
class MetaDescription extends React.Component {
  setMetaDescription = desc => {
    const metaList = document.getElementsByTagName('META');
    const metaDescription = [...metaList].filter(
      meta => meta.getAttribute('name') === 'description'
    )[0];
    metaDescription.setAttribute('content', desc);
  };

  componentDidMount() {
    this.setMetaDescription(this.props.children);
  }
  componentDidUpdate() {
    this.setMetaDescription(this.props.children);
  }
  render() {
    return null;
  }
}

MetaDescription.propTypes = {
  children: PropTypes.string.isRequired
};

export default MetaDescription;
