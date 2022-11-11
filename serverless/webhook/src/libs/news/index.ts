import { NewsClient } from '@hieudoanm/news';
import { API_KEY_NEWS } from '../../environments';

export const newsClient = new NewsClient({ apiKey: API_KEY_NEWS });

export default newsClient;
