import React from 'react';
import ScrollLink from './ScrollLink';

/** Navbar component is a navbar with MDL specific classes. */
const Navbar = () => {
  return (
    <header className="mdl-layout__header">
      <div className="mdl-layout__header-row">
        <span id="top-bar-title" className="mdl-layout-title">
          <ScrollLink
            style={{ textDecoration: 'none', color: 'white' }}
            to={'/'}
          >
            MathFever
          </ScrollLink>
        </span>
        <div className="mdl-layout-spacer" />

        <nav className="mdl-navigation mdl-layout--large-screen-only">
          <ScrollLink className={'mdl-navigation__link'} to={'/'}>
            Home
          </ScrollLink>
          <ScrollLink className={'mdl-navigation__link'} to={'/message-board'}>
            Message Board
          </ScrollLink>
          <ScrollLink className={'mdl-navigation__link'} to={'/about'}>
            About
          </ScrollLink>
          <ScrollLink className={'mdl-navigation__link'} to={'/help'}>
            Help
          </ScrollLink>
        </nav>
      </div>
    </header>
  );
};

export default Navbar;
