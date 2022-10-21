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

export type CoinsRequest = {
  limit?: number;
  offset?: number;
  orderBy?: OrderBy;
  orderDirection?: OrderDirection;
  tags?: Tag | string;
  timePeriod?: TimePeriod;
  tier?: Tier;
};
