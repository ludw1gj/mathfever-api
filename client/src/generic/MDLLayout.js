import React from 'react';

/**
 * MDLLayout is a div of MDL specific classes which childs the Navbar, Drawer, Content,
 * and Footer.
 */
const MDLLayout = ({ children }) => {
  return (
    <div className="mdl-layout mdl-js-layout mdl-layout--fixed-header">
      {children}
    </div>
  );
};

export default MDLLayout;
