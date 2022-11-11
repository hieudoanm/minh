import axios from '@hieudoanm/axios';
import logger from '@hieudoanm/pino';
import { chartify } from '../../libs/chartify';
import {
  HEIGHT,
  MAX_VALUE,
  STROKE_COLOR,
  STROKE_WIDTH,
  WIDTH,
} from '../../libs/chartify/defaults';
import { StockHistory } from './vnindex.types';

const TABLEBASE_URL =
  'https://raw.githubusercontent.com/hieudoanm/tablebase/master';

export const chartifyHistory = async (
  symbol: string,
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
    `${TABLEBASE_URL}/json/vietnam/stock/history/${symbol.toUpperCase()}.json`
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
