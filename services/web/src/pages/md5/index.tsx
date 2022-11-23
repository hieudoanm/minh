import { copyToClipboard } from '@hieudoanm/utils';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import TextField from '@mui/material/TextField';
import md5 from 'md5';
import { NextPage } from 'next';
import { FormEvent, useState } from 'react';

const MD5Page: NextPage = () => {
  const [value, setValue] = useState('');
  const [hash, setHash] = useState('MD5');

  return (
    <Container className="h-screen">
      <div className="h-full px-4">
        <div className="flex h-full items-center justify-center">
          <form className="flex flex-col border p-8 rounded gap-y-4">
            <TextField
              id="value"
              label="Value"
              placeholder="Value"
              value={value}
              className="w-full text-center"
              size="small"
              onChange={(event) => {
                const value = event.target.value;
                setValue(value);
                if (value === '') {
                  setHash('MD5');
                } else {
                  const hash = md5(value);
                  setHash(hash);
                }
              }}
            />
            <div
              className="py-2 px-4 border rounded cursor-pointer text-center"
              onClick={() => {
                copyToClipboard(hash);
                alert('Copy to Clipboard');
              }}
            >
              {hash}
            </div>
          </form>
        </div>
      </div>
    </Container>
  );
};

export default MD5Page;
