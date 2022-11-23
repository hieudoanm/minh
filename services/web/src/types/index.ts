export type CoinsStats = {
  total: number;
  totalCoins: number;
  totalMarkets: number;
  totalExchanges: number;
  totalMarketCap: string;
  total24hVolume: string;
};

export type Coin = {
  uuid: string;
  symbol: string;
  name: string;
  color: string;
  iconUrl: string;
  marketCap: string;
  price: string;
  listedAt: number;
  tier: number;
  change: string;
  rank: number;
  sparkline: string[];
  lowVolume: boolean;
  coinrankingUrl: string;
  btcPrice: string;
};

export type CoinsData = { stats: CoinsStats; coins: Coin[] };

export type CoinsResponse = { status: string; data: CoinsData };

export enum Tag {
  'DAO' = 'dao',
  'DEFI' = 'defi',
  'DEX' = 'dex',
  'EXCHANGE' = 'exchange',
  'GAMING' = 'gaming',
  'LAYER_1' = 'layer-1',
  'LAYER_2' = 'layer-2',
  'MEME' = 'meme',
  'METAVERSE' = 'metaverse',
  'NFT' = 'nft',
  'PRIVACY' = 'privacy',
  'STABLECOIN' = 'stablecoin',
  'STAKING' = 'staking',
  'WRAPPED' = 'wrapped',
}

export type ForexResponse = {
  success: boolean;
  timestamp: number;
  base: string;
  date: string;
  rates: Record<string, number>;
};

export type ForexRate = {
  code: string;
  currency: string;
  rate: number;
};
