import { NODE_ENV } from '../environments';

export const APP_NAME = 'MINH';

const isProduction: boolean = NODE_ENV === 'production';
export const PROD_BASE_URL = 'https://chatbot-functions.vercel.app';
export const DEV_BASE_URL = 'http://localhost:8080';
export const BASE_URL = isProduction ? PROD_BASE_URL : DEV_BASE_URL;

export const BASE_PATH = '';
