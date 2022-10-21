import axios from '@hieudoanm/axios';
import { getCoins } from '../crypto.service';

describe('crypto service', () => {
  describe('getCoins', () => {
    it('should return visas', async () => {
      jest.spyOn(axios, 'get').mockResolvedValueOnce({ status: 'success' });
      const coinsReponse = await getCoins({});
      expect(coinsReponse).toEqual({ status: 'success' });
    });
  });
});
