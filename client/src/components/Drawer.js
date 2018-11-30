import React from 'react';
import ScrollLink from '../components/ScrollLink';

/** Drawer component is the drawer for navigation for smaller displays. */
const Drawer = () => {
  const onClick = () => {
    const drawer = document.getElementsByClassName('mdl-layout__drawer')[0];
    drawer.classList.remove('is-visible');
    drawer.setAttribute('aria-hidden', 'true');

    const obfuscator = document.getElementsByClassName(
      'mdl-layout__obfuscator'
    )[0];
    obfuscator.classList.remove('is-visible');
  };

  return (
    <div className="mdl-layout__drawer">
      <span className="mdl-layout-title">MathFever</span>
      <nav className="mdl-navigation">
        <ScrollLink
          className={'mdl-navigation__link'}
          onClick={onClick}
          to={'/'}
        >
          Home
        </ScrollLink>
        <ScrollLink
          className={'mdl-navigation__link'}
          onClick={onClick}
          to={'/message-board'}
        >
          Message Board
        </ScrollLink>
        <ScrollLink
          className={'mdl-navigation__link'}
          onClick={onClick}
          to={'/about'}
        >
          About
        </ScrollLink>
        <ScrollLink
          className={'mdl-navigation__link'}
          onClick={onClick}
          to={'/help'}
        >
          Help
        </ScrollLink>
      </nav>
    </div>
  );
};

export default Drawer;
