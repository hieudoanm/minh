import logger from '@hieudoanm/pino';
import { telegramClient } from '../../libs/telegram';
import { weatherClient } from '../../libs/weather';
import { WebhookRequestBody } from './webhook.types';

const COMMANDS: string[] = ['weather'];

export const processWebhookRequestBody = async (
  requestBody: WebhookRequestBody
): Promise<void> => {
  logger.info('requestBody', requestBody);
  const {
    message: {
      chat: { id },
      text,
    },
  } = requestBody;
  const chatId: string = id.toString();
  const lowerCaseText = text.toLowerCase().trim();
  if (COMMANDS.includes(lowerCaseText)) {
    if (lowerCaseText === 'weather') {
      const message = await getWeatherMessage();
      await telegramClient.sendMessage(chatId, message);
    }
  } else {
    await telegramClient.sendMessage(chatId, 'N/A');
  }
};

const getWeatherMessage = async () => {
  const weather = await weatherClient.getWeather('ho chi minh city');
  const city = weather.name || '';
  const mainWeather: string = weather.weather[0].main || '';
  const description: string = weather.weather[0].description || '';
  const temp: number = weather.main.temp || 0;
  const feelsLike: number = weather.main.feels_like || 0;
  return `${city}\n${mainWeather} (${description})\n${temp}°C - ${feelsLike}°C`;
};
