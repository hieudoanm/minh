import axios from '@hieudoanm/axios';
import { getForex } from '../forex.service';
import { ForexResponse } from '../forex.types';

describe('forex service', () => {
  describe('getForex', () => {
    it('should return visas', async () => {
      jest.spyOn(axios, 'get').mockResolvedValueOnce({ status: 'success' });
      const forexResponse: ForexResponse = await getForex();
      expect(forexResponse).toEqual({ status: 'success' });
    });
  });
});
