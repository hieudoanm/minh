import {
  Controller,
  Example,
  Get,
  OperationId,
  Path,
  Query,
  Route,
  SuccessResponse,
  Tags,
} from '@hieudoanm/express';
import { coinsResponseExample } from './crypto.example';
import { getCoin, getCoins } from './crypto.service';
import {
  CoinResponse,
  CoinsResponse,
  OrderBy,
  OrderDirection,
  Tag,
  Tier,
  TimePeriod,
} from './crypto.types';

@Route('api/crypto')
@Tags('Crypto')
export class CryptoController extends Controller {
  @Get('coins')
  @OperationId('GetCryptoCoins')
  @SuccessResponse('200', 'List of Crypto Coins')
  @Example<CoinsResponse>(coinsResponseExample)
  public async getCoins(
    @Query('offset') offset = 0,
    @Query('limit') limit = 100,
    @Query('orderBy') orderBy: OrderBy = 'marketCap',
    @Query('orderDirection') orderDirection: OrderDirection = 'desc',
    @Query('timePeriod') timePeriod: TimePeriod = '24h',
    @Query('tier') tier: Tier = 1,
    @Query('tag') tag?: Tag
  ): Promise<CoinsResponse> {
    return getCoins({
      limit,
      offset,
      orderBy,
      orderDirection,
      timePeriod,
      tier,
      tag,
    });
  }

  @Get('coins/:id')
  @OperationId('GetCryptoCoin')
  @SuccessResponse('200', 'Get Crypto Coin')
  public async getCoin(@Path('id') id: string): Promise<CoinResponse> {
    return getCoin(id);
  }
}
