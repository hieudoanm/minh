import axios from '@hieudoanm/axios';
import logger from '@hieudoanm/pino';
import { API_KEY_COIN_RANKING } from '../../environments';
import { CoinResponse, CoinsRequest, CoinsResponse } from './crypto.types';

export const getCoins = async (
  {
    limit = 100,
    offset = 0,
    orderBy = 'marketCap',
    orderDirection = 'desc',
    timePeriod = '24h',
    tier = 1,
    tag,
  }: CoinsRequest = {
    limit: 100,
    offset: 0,
    orderBy: 'marketCap',
    orderDirection: 'desc',
    timePeriod: '24h',
    tier: 1,
  }
): Promise<CoinsResponse> => {
  const query = new URLSearchParams();
  // Pagination
  query.set('limit', limit.toString());
  query.set('offset', offset.toString());
  // Order
  query.set('orderBy', orderBy);
  query.set('orderDirection', orderDirection);
  // Categorize
  if (!tier) query.set('tier', tier);
  if (tag !== undefined) query.set('tags', tag);
  if (!timePeriod) query.set('timePeriod', timePeriod);
  // Call API
  const url = `https://api.coinranking.com/v2/coins?${query.toString()}`;
  const configs = { headers: { 'x-access-token': API_KEY_COIN_RANKING } };
  try {
    return await axios.get<CoinsResponse>(url, configs);
  } catch (error) {
    logger.error(`getCoins Error ${error}`);
    throw error;
  }
};

export const getCoin = async (id: string): Promise<CoinResponse> => {
  const url = `https://api.coinranking.com/v2/coin/${id}`;
  const configs = { headers: { 'x-access-token': API_KEY_COIN_RANKING } };
  try {
    return await axios.get<CoinResponse>(url, configs);
  } catch (error) {
    logger.error(`getCoin Error ${error}`);
    throw error;
  }
};
