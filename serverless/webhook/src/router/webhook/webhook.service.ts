import logger from '@hieudoanm/pino';
import { telegramClient } from '../../libs/telegram';
import { getNextMatchMessage } from '../../services/football.service';
import { getForexMessage } from '../../services/forex.service';
import {
  getBlockchainCryptoNewsMessage,
  getTopHeadlinesMessage,
} from '../../services/news.service';
import { getWeatherMessage } from '../../services/weather.service';
import { WebhookRequestBody } from './webhook.types';

const COMMANDS: string[] = [
  'blockchain news',
  'forex',
  'liverpool',
  'news',
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
    if (COMMANDS.includes(message)) {
      if (message === 'forex') {
        return await getForexMessage();
      } else if (message === 'weather') {
        return await getWeatherMessage();
      } else if (message === 'liverpool') {
        return await getNextMatchMessage();
      } else if (message === 'news') {
        return await getTopHeadlinesMessage();
      } else if (message === 'blockchain news') {
        return await getBlockchainCryptoNewsMessage();
      } else {
        return 'N/A';
      }
    } else if (message === 'help') {
      return COMMANDS.map((command: string) => `- \`${command}\``).join('\n');
    } else {
      return 'N/A';
    }
  } catch (error) {
    logger.error('Error', error);
    return 'Error';
  }
};
