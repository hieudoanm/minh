import { NextPage } from 'next';
import Map from '../../components/organisms/Map';

const MapPage: NextPage<{ countries: any[] }> = ({ countries = [] }) => {
  return (
    <div
      style={{
        width: '100vw',
        height: '100vh',
        overflow: 'hidden',
        backgroundColor: '#111827',
      }}
    >
      <Map countries={countries} />
    </div>
  );
};

export const getStaticProps = async () => {
  const response = await fetch('https://restcountries.com/v3/all');
  const countries = await response.json();
  return { props: { countries } };
};

export default MapPage;
