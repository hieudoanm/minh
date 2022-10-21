import { SuccessResponse } from '@hieudoanm/express';
import { Controller, Get, Query, Route, Tags } from 'tsoa';
import { Category, getForex } from './forex.service';
import { ForexResponse } from './forex.types';

@Route('api/forex')
@Tags('Forex')
export class ForexController extends Controller {
  @Get('rates')
  @SuccessResponse(200, 'List of Forex Rates')
  public async getForex(
    @Query('category') category: Category = null
  ): Promise<ForexResponse> {
    return getForex({ category });
  }
}
