export type StockHistory = {
  date: string;
  symbol: string;
  open: string;
  high: string;
  low: string;
  close: string;
  volume: string;
  timestamp: string;
};

export type StockCompany = {
  symbol: string;
  market: string;
  name: string;
  industry: string;
  supersector: string;
  sector: string;
  subsector: string;
  listedDate: string;
  issueShare: string;
  marketCap: string;
  priceChangedFiveDayPercent: string;
  priceChangedOneMonthPercent: string;
  priceChangedThreeMonthsPercent: string;
};
