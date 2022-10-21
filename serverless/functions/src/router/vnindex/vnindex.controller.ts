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
  WIDTH,
  HEIGHT,
  MAX_VALUE,
  STROKE_COLOR,
  STROKE_WIDTH,
} from '../../libs/chartify/defaults';
import { chartifyHistory, getCompanies, getHistory } from './vnindex.service';
import { StockCompany, StockHistory } from './vnindex.types';

@Tags('VNINDEX')
@Route('api/vnindex')
export class VnindexController extends Controller {
  @Get('companies')
  public async getCompanies(): Promise<StockCompany[]> {
    return getCompanies();
  }

  @Get('history/:stockCode')
  @SuccessResponse('200', 'Chart SVG')
  public async getHistory(
    @Path('stockCode') stockCode: string,
    @Query('offset') offset = 0,
    @Query('limit') limit = 20
  ): Promise<StockHistory[]> {
    return getHistory(stockCode, { offset, limit });
  }

  @Get('history/:stockCode/chart')
  @SuccessResponse('200', 'Chart SVG')
  public async chartifyHistory(
    @Request() request: ExpressRequest,
    @Path('stockCode') stockCode: string,
    @Query('width') width = WIDTH,
    @Query('height') height = HEIGHT,
    @Query('minValue') minValue = 0,
    @Query('maxValue') maxValue = MAX_VALUE,
    @Query('strokeColor') strokeColor = STROKE_COLOR,
    @Query('strokeWidth') strokeWidth = STROKE_WIDTH
  ): Promise<void> {
    const response: Response = request.res as Response;
    this.setStatus(200);
    this.setHeader('Content-Type', 'image/svg+xml');
    const chart = await chartifyHistory(stockCode, {
      width,
      height,
      minValue,
      maxValue,
      strokeColor,
      strokeWidth,
    });
    response.send(chart);
    return;
  }
}
