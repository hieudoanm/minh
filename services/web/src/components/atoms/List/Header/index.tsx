import React from 'react';

const Header: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return <div className="px-8 py-4">{children}</div>;
};

export default Header;
