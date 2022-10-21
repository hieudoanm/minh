import {
  Controller,
  Get,
  Path,
  Route,
  SuccessResponse,
  Tags,
} from '@hieudoanm/express';
import { SERVICE_NAMES } from './status.constants';
import { getStatus, getStatuses } from './status.service';
import { Service, Status, StatusesResponse } from './status.types';

@Tags('Status')
@Route('api/status')
export class StatusesController extends Controller {
  @Get()
  @SuccessResponse(200, 'Status of All Services')
  public async getStatuses(): Promise<StatusesResponse> {
    return getStatuses();
  }

  @Get(':service')
  @SuccessResponse(200, 'Status of the Service')
  public async getServiceStatus(
    @Path('service') service: Service
  ): Promise<{ name: string; status: Status }> {
    const status = await getStatus(service);
    const name = SERVICE_NAMES[service];
    return { name, status };
  }
}
