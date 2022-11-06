import logger from '@hieudoanm/pino';
import { v4 } from 'uuid';
import { telegramClient } from '../../libs/telegram';
import { getAirVisualMessage } from '../../services/air-visual.service';
import { getMatchesMessage } from '../../services/football.service';
import { getForexMessage } from '../../services/forex.service';
import {
  getBlockchainCryptoNewsMessage,
  getTopHeadlinesMessage,
} from '../../services/news.service';
import { getWeatherMessage } from '../../services/weather.service';
import { WebhookRequestBody } from './webhook.types';

const INTENTS: string[] = [
  'blockchain news',
  'forex',
  'liverpool',
  'news',
  'uuid',
  'weather',
].sort();

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

  const message: string = await processMessage(lowerCaseText);
  await telegramClient.sendMessage(chatId, message);
};

const processMessage = async (message: string): Promise<string> => {
  try {
    if (INTENTS.includes(message)) {
      if (message === 'forex') {
        return await getForexMessage();
      } else if (message === 'weather') {
        const airVisualMessage = await getAirVisualMessage();
        const weatherMessage = await getWeatherMessage();
        return `${weatherMessage}\n${airVisualMessage}`;
      } else if (message === 'liverpool') {
        return await getMatchesMessage();
      } else if (message === 'news') {
        return await getTopHeadlinesMessage();
      } else if (message === 'blockchain news') {
        return await getBlockchainCryptoNewsMessage();
      } else if (message === 'uuid') {
        return `\`${v4()}\``;
      } else {
        return 'N/A';
      }
    } else if (message === 'help') {
      return INTENTS.map((intent: string) => `- \`${intent}\``).join('\n');
    } else {
      return 'N/A';
    }
  } catch (error) {
    logger.error('Error', error);
    return 'Error';
  }
};
