import {
  Controller,
  Get,
  Path,
  Query,
  Request,
  Route,
  SuccessResponse,
  Tags,
} from '@hieudoanm/express';
import { Request as ExpressRequest, Response } from 'express';
import {
  HEIGHT,
  MAX_VALUE,
  STROKE_COLOR,
  STROKE_WIDTH,
  WIDTH,
} from '../../libs/chartify/defaults';
import { chartifyHistory } from './vnindex.service';

@Tags('VNINDEX')
@Route('api/vnindex')
export class VnindexController extends Controller {
  @Get('history/:symbol/chart')
  @SuccessResponse('200', 'Chart SVG')
  public async chartifyHistory(
    @Request() request: ExpressRequest,
    @Path('symbol') symbol: string,
    @Query('width') width = WIDTH,
    @Query('height') height = HEIGHT,
    @Query('minValue') minValue = 0,
    @Query('maxValue') maxValue = MAX_VALUE,
    @Query('strokeColor') strokeColor = STROKE_COLOR,
    @Query('strokeWidth') strokeWidth = STROKE_WIDTH
  ): Promise<null> {
    const response: Response = request.res as Response;
    this.setStatus(200);
    this.setHeader('Content-Type', 'image/svg+xml');
    const chart = await chartifyHistory(symbol, {
      width,
      height,
      minValue,
      maxValue,
      strokeColor,
      strokeWidth,
    });
    response.send(chart);
    return null;
  }
}
