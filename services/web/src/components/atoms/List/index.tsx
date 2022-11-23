import React from 'react';
import Header from './Header';
import Item from './Item';

const List: React.FC<{ className?: string; children: React.ReactNode }> = ({
  className = '',
  children,
}) => {
  return <div className={`border rounded ${className}`}>{children}</div>;
};

List.displayName = 'List';
List.defaultProps = { className: '' };

export default Object.assign(List, { Header, Item });
