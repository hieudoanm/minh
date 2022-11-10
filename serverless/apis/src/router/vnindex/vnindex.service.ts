import axios from '@hieudoanm/axios';
import logger from '@hieudoanm/pino';
import { chartify } from '../../libs/chartify';
import {
  WIDTH,
  HEIGHT,
  MAX_VALUE,
  STROKE_COLOR,
  STROKE_WIDTH,
} from '../../libs/chartify/defaults';
import { StockCompany, StockHistory } from './vnindex.types';

const TABLEBASE_URL =
  'https://raw.githubusercontent.com/hieudoanm/tablebase/master';

export const getCompanies = async (): Promise<StockCompany[]> => {
  return axios.get<StockCompany[]>(
    `${TABLEBASE_URL}/json/vietnam/stock/companies.json`
  );
};

export const getHistory = async (
  stockCode: string,
  { offset = 0, limit = 20 }: { offset: number; limit: number } = {
    offset: 0,
    limit: 20,
  }
): Promise<StockHistory[]> => {
  const history: StockHistory[] = await axios.get<StockHistory[]>(
    `${TABLEBASE_URL}/json/vietnam/stock/history/${stockCode.toUpperCase()}.json`
  );
  history.sort((a, b) => (a.date < b.date ? 1 : -1));
  return history.slice(offset, limit);
};

export const chartifyHistory = async (
  stockCode: string,
  {
    width = WIDTH,
    height = HEIGHT,
    minValue = 0,
    maxValue = MAX_VALUE,
    strokeColor = STROKE_COLOR,
    strokeWidth = STROKE_WIDTH,
  }: {
    width?: number;
    height?: number;
    minValue?: number;
    maxValue?: number;
    strokeColor?: string;
    strokeWidth?: number;
  } = {
    width: WIDTH,
    height: HEIGHT,
    minValue: 0,
    maxValue: MAX_VALUE,
    strokeColor: STROKE_COLOR,
    strokeWidth: STROKE_WIDTH,
  }
) => {
  const history: StockHistory[] = await axios.get<StockHistory[]>(
    `${TABLEBASE_URL}/json/vietnam/stock/history/${stockCode.toUpperCase()}.json`
  );
  history.sort((a, b) => (a.date < b.date ? 1 : -1));
  const data: number[] = history
    .map((value) => parseFloat(value.close))
    .slice(0, 20);
  logger.info('data', data);
  return chartify({
    data,
    scaleOptions: { minValue, maxValue },
    sizeOptions: { width, height },
    strokeOptions: { strokeColor, strokeWidth },
  });
};
