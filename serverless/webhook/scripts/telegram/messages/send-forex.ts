import dotenv from '@hieudoanm/dotenv';
dotenv.config();

import logger from '@hieudoanm/pino';
import { TELEGRAM_CHAT_ID } from '../../../src/environments';
import { telegramClient } from '../../../src/libs/telegram';
import { getForexMessage } from '../../../src/services/forex.service';

const main = async () => {
  try {
    const message = await getForexMessage();
    await telegramClient.sendMessage(TELEGRAM_CHAT_ID, message);
    process.exit(0);
  } catch (error) {
    logger.error('Error', error);
  }
};

main().catch((error: Error) => logger.error('Error', error));
