import { Body, Controller, Post, Route, Tags } from '@hieudoanm/express';
import { processWebhookRequestBody } from './webhook.service';
import { WebhookRequestBody } from './webhook.types';

@Tags('Webhook')
@Route('api/webhook')
export class WebhookController extends Controller {
  @Post()
  public async webhook(@Body() requestBody: WebhookRequestBody) {
    await processWebhookRequestBody(requestBody);
    return { status: 'OK' };
  }
}
