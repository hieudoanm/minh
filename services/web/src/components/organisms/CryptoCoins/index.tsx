import { useAxios } from '@hieudoanm/hooks';
import CircularProgress from '@mui/material/CircularProgress';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import { NextPage } from 'next';
import React, { ReactNode, useState } from 'react';
import { BASE_URL } from '../../../configs';
import { Coin, CoinsResponse, Tag } from '../../../types';
import List from '../../atoms/List';
import Crypto from '../../molecules/Crypto';

type CryptoParams = { tags: Tag | string; offset: number };

const CryptoCoinsList: React.FC<{ children: ReactNode }> = ({ children }) => {
  return (
    <List>
      <List.Header>
        <h1 className="text-2xl uppercase">Coins</h1>
      </List.Header>
      {children}
    </List>
  );
};

const CryptoCoins: NextPage = () => {
  const [params, setParams] = useState<CryptoParams>({ tags: '', offset: 0 });
  const url: string = `${BASE_URL}/api/crypto/coins`;
  const { data, loading, error, refetch } = useAxios<CoinsResponse>(url, {
    params,
  });

  if (loading) {
    return (
      <CryptoCoinsList>
        <List.Item className="flex items-center justify-center">
          <CircularProgress />
        </List.Item>
      </CryptoCoinsList>
    );
  }

  if (error) {
    const errorMessage: string = error.stack || error.message || 'Error';
    return (
      <CryptoCoinsList>
        <List.Item>
          <p className="p-4 text-center">{errorMessage}</p>
        </List.Item>
      </CryptoCoinsList>
    );
  }

  if (!data) {
    return (
      <CryptoCoinsList>
        <List.Item>
          <p>Response is null</p>
        </List.Item>
      </CryptoCoinsList>
    );
  }

  const coins: Coin[] = data.data.coins || [];
  const total: number = data.data.stats.total || 0;

  const maxPage: number =
    Math.floor(total / 100) <= 10 ? Math.floor(total / 100) : 10;
  const pages = [...Array(maxPage).keys()];

  return (
    <List>
      <List.Header>
        <div className="flex justify-between items-center">
          <h1 className="text-xl uppercase">Coins ({coins.length})</h1>
          <div className="flex gap-4">
            <FormControl size="small">
              <InputLabel id="label-tags">Tags</InputLabel>
              <Select
                labelId="label-tags"
                id="tags"
                value={params.tags}
                label="Tags"
                onChange={(event) => {
                  const tags = event.target.value as Tag;
                  setParams({ ...params, tags });
                  refetch();
                }}
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value={Tag.DAO}>
                  DAO
                  <span className="hidden sm:inline ml-1">
                    (Decentralized Autonomous Organization)
                  </span>
                </MenuItem>
                <MenuItem value={Tag.DEFI}>
                  DeFi
                  <span className="hidden sm:inline ml-1">
                    (Decentralized Finance)
                  </span>
                </MenuItem>
                <MenuItem value={Tag.DEX}>
                  DEX
                  <span className="hidden sm:inline ml-1">
                    (Decentralized Exchange)
                  </span>
                </MenuItem>
                <MenuItem value={Tag.EXCHANGE}>Exchange</MenuItem>
                <MenuItem value={Tag.GAMING}>Gaming</MenuItem>
                <MenuItem value={Tag.LAYER_1}>Layer 1</MenuItem>
                <MenuItem value={Tag.LAYER_2}>Layer 2</MenuItem>
                <MenuItem value={Tag.MEME}>Meme</MenuItem>
                <MenuItem value={Tag.METAVERSE}>Metaverse</MenuItem>
                <MenuItem value={Tag.NFT}>
                  NFT
                  <span className="hidden sm:inline ml-1">
                    (Non-fungible Token)
                  </span>
                </MenuItem>
                <MenuItem value={Tag.STABLECOIN}>Stable Coin</MenuItem>
                <MenuItem value={Tag.STAKING}>Staking</MenuItem>
                <MenuItem value={Tag.WRAPPED}>Wrapped</MenuItem>
              </Select>
            </FormControl>
            <FormControl size="small">
              <InputLabel id="label-page">Page</InputLabel>
              <Select
                labelId="label-page"
                id="page"
                value={params.tags}
                label="Page"
                onChange={(event) => {
                  const page: number = parseFloat(event.target.value);
                  const offset: number = page * 100;
                  setParams({ ...params, offset });
                  refetch();
                }}
              >
                {pages.map((page) => {
                  return (
                    <MenuItem key={page} value={page}>
                      {page + 1}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </div>
        </div>
      </List.Header>
      <Crypto coins={coins} />
    </List>
  );
};

export default CryptoCoins;
