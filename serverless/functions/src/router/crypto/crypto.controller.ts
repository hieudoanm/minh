import {
  Controller,
  Get,
  Query,
  Route,
  SuccessResponse,
  Tags,
} from '@hieudoanm/express';
import { Tag, TimePeriod, Tier, OrderBy, OrderDirection } from './coins.enums';
import { getCoins } from './crypto.service';
import { CoinsResponse } from './crypto.types';

@Route('api/crypto')
@Tags('Crypto')
export class CryptoController extends Controller {
  @Get('coins')
  @SuccessResponse('200', 'List of Coins')
  public async getCoins(
    @Query('limit') limit = 100,
    @Query('offset') offset = 0,
    @Query('orderBy') orderBy: OrderBy = OrderBy.MARKET_CAP,
    @Query('orderDirection')
    orderDirection: OrderDirection = OrderDirection.DESC,
    @Query('tags') tags: Tag | string = '',
    @Query('timePeriod') timePeriod: TimePeriod = TimePeriod.HOUR_24,
    @Query('tier') tier: Tier = Tier.TIER_1
  ): Promise<CoinsResponse> {
    return getCoins({
      limit,
      offset,
      orderBy,
      orderDirection,
      tags,
      timePeriod,
      tier,
    });
  }
}
