import React, { Component } from 'react';

/** mdlUpgrade adds firing MDL componentHandler.upgradeDom function on component mount. */
const mdlUpgrade = WrappedComponent => {
  return class MDLUpgrade extends Component {
    componentDidMount() {
      window.componentHandler.upgradeDom();
    }
    render() {
      return <WrappedComponent {...this.props} />;
    }
  };
};

export default mdlUpgrade;
