import React from 'react';
import styled from 'styled-components';

import ScrollLink from '../components/ScrollLink';
import SmallCaps from './SmallCaps';

const Wrapper = styled.footer`
  margin-top: 50px;
`;

/** Footer component is a footer with MDL specific classes. */
const Footer = () => {
  return (
    <Wrapper className="mdl-mini-footer">
      <div className="mdl-mini-footer__left-section">
        <div className="mdl-logo">
          <SmallCaps>&#169; 2018 MathFever</SmallCaps>
        </div>
        <ul className="mdl-mini-footer__link-list">
          <li>
            <ScrollLink to={'/help'}>Help</ScrollLink>
          </li>
          <li>
            <ScrollLink to={'/terms'}>Terms of Use</ScrollLink>
          </li>
          <li>
            <ScrollLink to={'/privacy'}>Privacy</ScrollLink>
          </li>
        </ul>
      </div>
    </Wrapper>
  );
};

export default Footer;
