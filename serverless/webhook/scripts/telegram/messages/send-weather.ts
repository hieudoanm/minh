import dotenv from '@hieudoanm/dotenv';
dotenv.config();

import logger from '@hieudoanm/pino';
import { TELEGRAM_CHAT_ID } from '../../../src/environments';
import telegramClient from '../../../src/libs/telegram';
import { getAirVisualMessage } from '../../../src/services/air-visual.service';
import { getWeatherMessage } from '../../../src/services/weather.service';

const main = async () => {
  const airVisualMessage = await getAirVisualMessage();
  const weatherMessage = await getWeatherMessage();
  const message = `${weatherMessage}\n${airVisualMessage}`;
  await telegramClient.sendMessage(TELEGRAM_CHAT_ID, message);
  process.exit(0);
};

main().catch((error: Error) => logger.error('Error', error));
