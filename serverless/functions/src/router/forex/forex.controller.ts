import { OperationId, SuccessResponse } from '@hieudoanm/express';
import { Controller, Get, Query, Route, Tags } from 'tsoa';
import { Category, getForexRates } from './forex.service';
import { ForexResponse } from './forex.types';

@Route('api/forex')
@Tags('Forex')
export class ForexController extends Controller {
  @Get('rates')
  @OperationId('GetForexRates')
  @SuccessResponse(200, 'List of Forex Rates')
  public async getForexRates(
    @Query('category') category: Category = null
  ): Promise<ForexResponse> {
    return getForexRates({ category });
  }
}
