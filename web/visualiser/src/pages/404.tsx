import Button from '@mui/material/Button';
import { NextPage } from 'next';
import Link from 'next/link';

const NotFoundPage: NextPage = () => {
  return (
    <div className="w-screen h-screen">
      <div className="w-full h-full flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-6xl">404</h1>
          <div className="mt-4">
            <Link href="/">
              <Button variant="outlined" size="large">
                Go Home
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default NotFoundPage;
