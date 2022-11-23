import Container from '@mui/material/Container';
import type { NextPage } from 'next';
import CryptoCoins from '../../../components/organisms/CryptoCoins';

const CryptoPage: NextPage = () => {
  return (
    <Container className="py-4 md:py-8">
      <CryptoCoins />
    </Container>
  );
};

export default CryptoPage;
