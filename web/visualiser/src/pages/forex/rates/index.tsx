import Container from '@mui/material/Container';
import type { NextPage } from 'next';
import ForexRates from '../../../components/organisms/ForexRates';

const ForexPage: NextPage = () => {
  return (
    <Container className="py-4 md:py-8">
      <ForexRates />
    </Container>
  );
};

export default ForexPage;
