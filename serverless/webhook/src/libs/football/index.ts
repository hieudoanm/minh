import { FootballClient } from '@hieudoanm/football';
import { API_KEY_FOOTBALL_DATA } from '../../environments';

export const footballClient = new FootballClient({
  apiKey: API_KEY_FOOTBALL_DATA,
});

export default footballClient;
