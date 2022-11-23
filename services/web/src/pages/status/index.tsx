import axios from '@hieudoanm/axios';
import Container from '@mui/material/Container';
import type { NextPage } from 'next';
import List from '../../components/atoms/List';
import { BASE_URL } from '../../configs';

type Status = { id: string; service: string; status: string };

const StatusPage: NextPage<{ statuses: Status[] }> = ({ statuses = [] }) => {
  return (
    <Container className="py-4 md:py-8">
      <List>
        <List.Header>
          <div className="flex justify-between items-center">
            <h1 className="uppercase text-2xl">Service</h1>
            <h1 className="uppercase text-2xl">Status</h1>
          </div>
        </List.Header>
        {statuses.map((value) => {
          const { service, id, status } = value;
          const bgColor = status === 'active' ? 'bg-green-500' : 'bg-red-500';
          return (
            <List.Item key={`status-${id}`}>
              <div className="flex justify-between items-center">
                <div>{service}</div>
                <div className="px-4 py-2 rounded border rounded-full">
                  <div className="flex items-center gap-2">
                    <div className={`rounded-full w-4 h-4 ${bgColor}`} />
                    <p className="text-sm leading-4 uppercase">{status}</p>
                  </div>
                </div>
              </div>
            </List.Item>
          );
        })}
      </List>
    </Container>
  );
};

export const getStaticProps = async () => {
  const url: string = `${BASE_URL}/api/status`;
  const status = await axios.get<Record<string, object>>(url);
  const statuses = Object.keys(status).map((service: string) => {
    const value = status[service];
    return { ...value, id: status };
  });
  console.log(status);
  return { props: { statuses } };
};

export default StatusPage;
