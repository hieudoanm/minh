import express, { notFoundHandler, errorHandler } from '@hieudoanm/express';
import { RegisterRoutes } from './routes';

const app = express();

RegisterRoutes(app);

app.use(notFoundHandler());
app.use(errorHandler);

export default app;
