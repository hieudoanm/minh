import dotenv from '@hieudoanm/dotenv';
dotenv.config();

import logger from '@hieudoanm/pino';
import { TELEGRAM_CHAT_ID } from '../../../src/environments';
import { telegramClient } from '../../../src/libs/telegram';
import {
  getBlockchainCryptoNewsMessage,
  getTopHeadlinesMessage,
} from '../../../src/services/news.service';

const main = async () => {
  try {
    const topHeadlinesMessage = await getTopHeadlinesMessage();
    await telegramClient.sendMessage(TELEGRAM_CHAT_ID, topHeadlinesMessage);
    const blockchainNewsMessage = await getBlockchainCryptoNewsMessage();
    await telegramClient.sendMessage(TELEGRAM_CHAT_ID, blockchainNewsMessage);
    process.exit(0);
  } catch (error) {
    logger.error('Error', error);
    process.exit(1);
  }
};

main().catch((error: Error) => logger.error('Error', error));
