import { TelegramClient } from '@hieudoanm/telegram';
import { TELEGRAM_TOKEN } from '../../environments';

export const telegramClient = new TelegramClient({ token: TELEGRAM_TOKEN });

export default telegramClient;
