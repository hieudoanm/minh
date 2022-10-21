import dotenv from '@hieudoanm/dotenv';
dotenv.config();

import logger from '@hieudoanm/pino';
import { TELEGRAM_CHAT_ID } from '../../../src/environments';
import { telegramClient } from '../../../src/libs/telegram';
import { getMatchesMessage } from '../../../src/services/football.service';

const main = async () => {
  const message = await getMatchesMessage();
  await telegramClient.sendMessage(TELEGRAM_CHAT_ID, message);
  process.exit(0);
};

main().catch((error: Error) => logger.error('Error', error));
