import axios from '@hieudoanm/axios';
import { CryptoController } from '../crypto.controller';

describe('crypto controller', () => {
  const cryptoController = new CryptoController();

  describe('getCoins', () => {
    it('should return visas', async () => {
      jest.spyOn(axios, 'get').mockResolvedValueOnce({ status: 'success' });
      const coinsReponse = await cryptoController.getCoins();
      expect(coinsReponse).toEqual({ status: 'success' });
    });
  });
});
