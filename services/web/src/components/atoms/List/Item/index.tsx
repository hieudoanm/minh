import React from 'react';

const Item: React.FC<{ className?: string; children: React.ReactNode }> = ({
  className,
  children,
}) => {
  return <div className={`px-8 py-4 border-t ${className}`}>{children}</div>;
};

Item.displayName = 'List.Item';
Item.defaultProps = { className: '' };

export default Item;
