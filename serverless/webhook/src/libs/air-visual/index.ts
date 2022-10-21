import { AirVisualClient } from '@hieudoanm/air-visual';
import { API_KEY_AIR_VISUAL } from '../../environments';

const airVisualClient = new AirVisualClient({ apiKey: API_KEY_AIR_VISUAL });

export default airVisualClient;
