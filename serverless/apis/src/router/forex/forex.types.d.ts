export type ForexResponse = {
  success: boolean;
  timestamp: number;
  base: string;
  date: string;
  rates: Record<string, number>;
};
