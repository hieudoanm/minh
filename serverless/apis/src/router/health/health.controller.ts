import { Example, OperationId, SuccessResponse } from '@hieudoanm/express';
import { Controller, Get, Route, Tags } from 'tsoa';

@Tags('Health')
@Route('api/health')
export class HealthController extends Controller {
  @Get()
  @OperationId('GetHealth')
  @SuccessResponse('200', 'Get Health of Service')
  @Example<{ status: string }>({ status: 'healthy' })
  public getHealth() {
    return { status: 'healthy' };
  }
}
