import React from 'react';

/** MDLContent is a div of MDL specific classes which childs the Content and Footer. */
const MDLContent = ({ children }) => {
  return <main className="mdl-layout__content">{children}</main>;
};

export default MDLContent;
