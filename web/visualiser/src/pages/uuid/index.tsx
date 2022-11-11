import { copyToClipboard } from '@hieudoanm/utils';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import TextField from '@mui/material/TextField';
import { NextPage } from 'next';
import { FormEvent, useState } from 'react';
import { v4 } from 'uuid';

type UUIDPageProps = { uuid: string };

const UUIDPage: NextPage<UUIDPageProps> = ({ uuid }) => {
  const [id, setId] = useState(uuid);

  const generate = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setId(v4());
  };

  return (
    <Container>
      <div className="w-full h-screen">
        <div className="flex w-full h-full items-center justify-center">
          <form onSubmit={generate}>
            <div className="flex border p-8 rounded gap-4">
              <div
                className="py-2 px-4 border rounded cursor-pointer"
                onClick={() => {
                  copyToClipboard(id);
                  alert('Copy to Clipboard');
                }}
              >
                {id}
              </div>
              <Button type="submit" variant="outlined">
                NEW
              </Button>
            </div>
          </form>
        </div>
      </div>
    </Container>
  );
};

export const getStaticProps = () => {
  const uuid = v4();
  return { props: { uuid } };
};

export default UUIDPage;
