import { WeatherClient } from '@hieudoanm/weather';
import { API_KEY_OPEN_WEATHER_MAP } from '../../environments';

export const weatherClient = new WeatherClient({
  apiKey: API_KEY_OPEN_WEATHER_MAP,
});

export default weatherClient;
