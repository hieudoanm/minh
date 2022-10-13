import { ForexClient } from '@hieudoanm/forex';
import { API_KEY_FIXER } from '../../environments';

export const forexClient = new ForexClient(API_KEY_FIXER);

export default forexClient;
