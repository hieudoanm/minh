import express, { errorHandler, notFoundHandler } from '@hieudoanm/express';
import logger from '@hieudoanm/pino';
import { Request, Response } from 'express';
import createHttpError from 'http-errors';
import request from 'request';
import { RegisterRoutes } from './routes';

const app = express();

RegisterRoutes(app);

app.get('/api/proxy', (exRequest: Request, response: Response) => {
  const url = exRequest.query.url?.toString();
  logger.info('url', { url });
  if (!url) {
    throw createHttpError(400, 'url is required');
  }
  exRequest.pipe(request(url)).pipe(response);
});

app.use(notFoundHandler());
app.use(errorHandler);

export default app;