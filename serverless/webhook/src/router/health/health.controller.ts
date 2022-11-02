import { Controller, Get, Route, Tags } from 'tsoa';

@Tags('Health')
@Route('api/health')
export class HealthController extends Controller {
  @Get()
  public get() {
    return { status: 'healthy' };
  }
}
