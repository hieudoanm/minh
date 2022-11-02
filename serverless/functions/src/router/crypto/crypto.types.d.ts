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
  '24hVolume': string;
  btcPrice: string;
};

export type CoinsData = { stats: CoinsStats; coins: Coin[] };

export type CoinsResponse = { status: string; data: CoinsData };

export type CoinsRequest = {
  limit?: number;
  offset?: number;
  orderBy?: OrderBy;
  orderDirection?: string;
  timePeriod?: TimePeriod;
  tier?: Tier;
  tag?: Tag;
};

export type Link = {
  name: string;
  type: string;
  url: string;
};

export type CoinResponse = {
  status: string;
  data: {
    coin: {
      uuid: string;
      symbol: string;
      name: string;
      description: string;
      color: string;
      iconUrl: string;
      websiteUrl: string;
      links: Link[];
      supply: {
        confirmed: boolean;
        supplyAt: number;
        max: string;
        total: string;
        circulating: string;
      };
      numberOfMarkets: number;
      numberOfExchanges: number;
      '24hVolume': string;
      marketCap: string;
      fullyDilutedMarketCap: string;
      price: string;
      btcPrice: string;
      priceAt: number;
      change: string;
      rank: number;
      sparkline: string[];
      allTimeHigh: {
        price: string;
        timestamp: number;
      };
      coinrankingUrl: string;
      tier: number;
      lowVolume: boolean;
      listedAt: number;
      tags: string[];
    };
  };
};

export type Tier = 1 | 2 | 3;
export type OrderDirection = 'asc' | 'desc';
export type OrderBy =
  | '24hVolume'
  | 'change'
  | 'listedAt'
  | 'marketCap'
  | 'price';
export type TimePeriod =
  | '1h'
  | '3h'
  | '12h'
  | '24h'
  | '7d'
  | '30d'
  | '3m'
  | '1y'
  | '3y'
  | '5y';
export type Tag =
  | 'dao'
  | 'defi'
  | 'dex'
  | 'exchange'
  | 'gaming'
  | 'layer-1'
  | 'layer-2'
  | 'meme'
  | 'metaverse'
  | 'nft'
  | 'privacy'
  | 'stablecoin'
  | 'staking'
  | 'wrapped';
