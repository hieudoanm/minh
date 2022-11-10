import axios from '@hieudoanm/axios';
import { SERVICES, SERVICE_NAMES, SERVICE_URLS } from './status.constants';
import {
  Service,
  ServiceStatusResponse,
  Status,
  StatusesResponse,
  StatusResponse,
} from './status.types';

export const getStatus = async (service: Service): Promise<Status> => {
  const url: string = SERVICE_URLS[service];
  const data = await axios.get<ServiceStatusResponse>(url);
  return data.status.indicator === 'none' ? 'active' : 'inactive';
};

export const getStatuses = async (): Promise<StatusesResponse> => {
  return new Promise((resolve) => {
    Promise.all(
      SERVICES.map(async (service) => {
        const status = await getStatus(service);
        return { service, status };
      })
    ).then((data: StatusResponse[]) => {
      const response: StatusesResponse = {} as StatusesResponse;
      for (const service of SERVICES) {
        const serviceName = SERVICE_NAMES[service];
        const { status = 'inactive' } =
          data.find((item) => item.service === service) ||
          ({} as StatusResponse);
        response[service] = { service: serviceName, status };
      }
      return resolve(response);
    });
  });
};
