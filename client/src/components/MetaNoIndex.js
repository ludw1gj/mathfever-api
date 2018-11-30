import React from 'react';

/** MetaNoIndex component sets the document noindex meta tag. */
class MetaNoIndex extends React.Component {
  componentDidMount() {
    const metaElement = document.createElement('meta');
    metaElement.name = 'robots';
    metaElement.content = 'noindex';
    document.head.appendChild(metaElement);
  }

  componentWillUnmount() {
    const metaList = document.getElementsByTagName('META');
    const metaElement = [...metaList].filter(
      meta => meta.getAttribute('name') === 'robots'
    )[0];
    document.head.removeChild(metaElement);
  }

  render() {
    return null;
  }
}

export default MetaNoIndex;
