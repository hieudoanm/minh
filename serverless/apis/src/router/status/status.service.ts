import axios from '@hieudoanm/axios';
import { SERVICE_URLS } from './status.constants';
import { Service, ServiceStatusResponse, Status } from './status.types';

export const getStatus = async (service: Service): Promise<Status> => {
  const url: string = SERVICE_URLS[service];
  const data = await axios.get<ServiceStatusResponse>(url);
  return data.status.indicator === 'none' ? 'active' : 'inactive';
};
