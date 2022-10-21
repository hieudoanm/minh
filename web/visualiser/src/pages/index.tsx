import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import type { NextPage } from 'next';
import Link from '../components/atoms/Link';

type App = {
  id: string;
  href: string;
  name: string;
};

const APPS: App[] = [
  { id: 'crypto-coins', href: '/crypto/coins', name: 'Crypto Coins' },
  { id: 'forex-rates', href: '/forex/rates', name: 'Forex Rates' },
  { id: 'pomodoro', href: '/pomodoro', name: 'Pomodoro' },
  { id: 'md5', href: '/md5', name: 'MD5' },
  { id: 'status', href: '/status', name: 'Status' },
  { id: 'uuid', href: '/uuid', name: 'UUID' },
];

const Home: NextPage = () => {
  return (
    <Container className="py-4 md:py-8">
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4 md:gap-8">
        {APPS.map((app: App) => {
          const { id, href, name } = app;
          return (
            <Link key={`app-${id}`} href={href}>
              <Button type="button" variant="outlined">
                {name}
              </Button>
            </Link>
          );
        })}
      </div>
    </Container>
  );
};

export default Home;
