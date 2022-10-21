import axios from '@hieudoanm/axios';
import logger from '@hieudoanm/pino';
import { API_KEY_COIN_RANKING } from '../../environments';
import { OrderBy, OrderDirection, Tier, TimePeriod } from './coins.enums';
import { CoinsRequest, CoinsResponse } from './crypto.types';

export const getCoins = async (
  {
    limit = 100,
    offset = 0,
    orderBy = OrderBy.MARKET_CAP,
    orderDirection = OrderDirection.DESC,
    tags = '',
    timePeriod = TimePeriod.HOUR_24,
    tier = Tier.TIER_1,
  }: CoinsRequest = {
    limit: 100,
    offset: 0,
    orderBy: OrderBy.MARKET_CAP,
    orderDirection: OrderDirection.DESC,
    tags: '',
    timePeriod: TimePeriod.HOUR_24,
    tier: Tier.TIER_1,
  }
): Promise<CoinsResponse> => {
  const query = new URLSearchParams();
  query.set('limit', limit.toString());
  query.set('offset', offset.toString());
  query.set('orderBy', orderBy);
  query.set('orderDirection', orderDirection);
  if (tags !== '') query.set('tags', tags);
  if (timePeriod !== '') query.set('timePeriod', timePeriod);
  if (tier !== '') query.set('tier', tier);
  const url = `https://api.coinranking.com/v2/coins?${query.toString()}`;
  const configs = { headers: { 'x-access-token': API_KEY_COIN_RANKING } };
  try {
    return await axios.get<CoinsResponse>(url, configs);
  } catch (error) {
    logger.error(`getCoins Error ${error}`);
    throw error;
  }
};
