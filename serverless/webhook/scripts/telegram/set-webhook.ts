import dotenv from '@hieudoanm/dotenv';
dotenv.config();

import logger from '@hieudoanm/pino';
import { TELEGRAM_WEBHOOK } from '../../src/environments';
import { telegramClient } from '../../src/libs/telegram';

const main = async () => {
  logger.info('TELEGRAM_WEBHOOK', { TELEGRAM_WEBHOOK });
  const deleteResponse = await telegramClient.deleteWebhook(TELEGRAM_WEBHOOK);
  logger.info('deleteResponse', { deleteResponse });
  const setResponse = await telegramClient.setWebhook(TELEGRAM_WEBHOOK);
  logger.info('setResponse', { setResponse });
  const webhookInfo = await telegramClient.getWebhookInfo();
  logger.info('webhookInfo', { webhookInfo });
};

main().catch((error) => logger.error('Error', error));
