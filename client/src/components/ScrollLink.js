import React from 'react';
import { Link } from 'react-router-dom';

import { scrollToTop } from '../util/scroll';

/**
 * ScrollLink component wraps the react-router-dom Link component, and enables scrolling
 * to the top of the MDL layout content on click.
 */
const ScrollLink = props => {
  const onClick = () => {
    // execute onClock if passed from props.
    if (props.onClick) {
      props.onClick();
    }

    // scroll to the top of the mdl layout content.
    const mdlLayoutContent = document.getElementsByClassName(
      'mdl-layout__content'
    )[0];
    scrollToTop(mdlLayoutContent, 100);
  };
  return <Link {...props} onClick={onClick} />;
};

export default ScrollLink;
