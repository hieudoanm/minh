import { weatherClient } from '../libs/weather';

export const getWeatherMessage = async (): Promise<string> => {
  const weather = await weatherClient.getWeather('ho chi minh city');
  const city = weather.name || '';
  const mainWeather: string = weather.weather[0].main || '';
  const description: string = weather.weather[0].description || '';
  const temp: number = weather.main.temp || 0;
  const feelsLike: number = weather.main.feels_like || 0;
  return `${city}\n${mainWeather} (${description})\n${temp}°C - ${feelsLike}°C`;
};
