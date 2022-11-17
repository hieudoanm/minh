import Link from 'next/link';
import { FC, ReactNode } from 'react';
import { BASE_PATH } from '../../../configs';

type CustomLinkProps = { href: string; children: ReactNode };

const CustomLink: FC<CustomLinkProps> = ({ href, children }) => {
  return <Link href={`${BASE_PATH}${href}`}>{children}</Link>;
};

export default CustomLink;
