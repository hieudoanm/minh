import React from 'react';

export const Navbar: React.FC = () => {
  return (
    <nav>
      <div className="border-b">
        <div className="container mx-auto px-8 py-4">
          <h1 className="text-2xl">MINI APPS</h1>
        </div>
      </div>
    </nav>
  );
};

Navbar.displayName = 'Navbar';
Navbar.defaultProps = {};

export default Navbar;
