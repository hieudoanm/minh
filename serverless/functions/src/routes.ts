/* tslint:disable */
/* eslint-disable */
// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
import {
  Controller,
  ValidationService,
  FieldErrors,
  ValidateError,
  TsoaRoute,
  HttpStatusCodeLiteral,
  TsoaResponse,
  fetchMiddlewares,
} from '@tsoa/runtime';
// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
import { CryptoController } from './router/crypto/crypto.controller';
// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
import { ForexController } from './router/forex/forex.controller';
// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
import { HealthController } from './router/health/health.controller';
// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
import { StatusesController } from './router/status/status.controller';
// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
import { VnindexController } from './router/vnindex/vnindex.controller';
import type { RequestHandler } from 'express';
import * as express from 'express';

// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

const models: TsoaRoute.Models = {
  CoinsStats: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        total24hVolume: { dataType: 'string', required: true },
        totalMarketCap: { dataType: 'string', required: true },
        totalExchanges: { dataType: 'double', required: true },
        totalMarkets: { dataType: 'double', required: true },
        totalCoins: { dataType: 'double', required: true },
        total: { dataType: 'double', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  Coin: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        btcPrice: { dataType: 'string', required: true },
        coinrankingUrl: { dataType: 'string', required: true },
        lowVolume: { dataType: 'boolean', required: true },
        sparkline: {
          dataType: 'array',
          array: { dataType: 'string' },
          required: true,
        },
        rank: { dataType: 'double', required: true },
        change: { dataType: 'string', required: true },
        tier: { dataType: 'double', required: true },
        listedAt: { dataType: 'double', required: true },
        price: { dataType: 'string', required: true },
        marketCap: { dataType: 'string', required: true },
        iconUrl: { dataType: 'string', required: true },
        color: { dataType: 'string', required: true },
        name: { dataType: 'string', required: true },
        symbol: { dataType: 'string', required: true },
        uuid: { dataType: 'string', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  CoinsData: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        coins: {
          dataType: 'array',
          array: { dataType: 'refAlias', ref: 'Coin' },
          required: true,
        },
        stats: { ref: 'CoinsStats', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  CoinsResponse: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        data: { ref: 'CoinsData', required: true },
        status: { dataType: 'string', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  OrderBy: {
    dataType: 'refEnum',
    enums: ['24hVolume', 'change', 'listedAt', 'marketCap', 'price'],
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  OrderDirection: {
    dataType: 'refEnum',
    enums: ['asc', 'desc'],
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  Tag: {
    dataType: 'refEnum',
    enums: [
      'dao',
      'defi',
      'dex',
      'exchange',
      'gaming',
      'layer-1',
      'layer-2',
      'meme',
      'metaverse',
      'nft',
      'privacy',
      'stablecoin',
      'staking',
      'wrapped',
    ],
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  TimePeriod: {
    dataType: 'refEnum',
    enums: ['1h', '3h', '12h', '24h', '7d', '30d', '3m', '1y', '3y', '5y'],
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  Tier: {
    dataType: 'refEnum',
    enums: ['1', '2', '3'],
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  'Record_string.number_': {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {},
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  ForexResponse: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        rates: { ref: 'Record_string.number_', required: true },
        date: { dataType: 'string', required: true },
        base: { dataType: 'string', required: true },
        timestamp: { dataType: 'double', required: true },
        success: { dataType: 'boolean', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  Category: {
    dataType: 'refAlias',
    type: {
      dataType: 'union',
      subSchemas: [
        { dataType: 'enum', enums: ['top'] },
        { dataType: 'enum', enums: ['personal'] },
        { dataType: 'enum', enums: [null] },
      ],
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  Status: {
    dataType: 'refAlias',
    type: {
      dataType: 'union',
      subSchemas: [
        { dataType: 'enum', enums: ['active'] },
        { dataType: 'enum', enums: ['inactive'] },
      ],
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  StatusResponse: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        status: { ref: 'Status', required: true },
        service: { dataType: 'string', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  'Record_Service.StatusResponse_': {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        atlassian: { ref: 'StatusResponse' },
        bitbucket: { ref: 'StatusResponse' },
        circleci: { ref: 'StatusResponse' },
        confluence: { ref: 'StatusResponse' },
        discord: { ref: 'StatusResponse' },
        github: { ref: 'StatusResponse' },
        hedera: { ref: 'StatusResponse' },
        'jira-software': { ref: 'StatusResponse' },
        solana: { ref: 'StatusResponse' },
        trello: { ref: 'StatusResponse' },
        vercel: { ref: 'StatusResponse' },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  StatusesResponse: {
    dataType: 'refAlias',
    type: { ref: 'Record_Service.StatusResponse_', validators: {} },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  Service: {
    dataType: 'refAlias',
    type: {
      dataType: 'union',
      subSchemas: [
        { dataType: 'enum', enums: ['atlassian'] },
        { dataType: 'enum', enums: ['bitbucket'] },
        { dataType: 'enum', enums: ['circleci'] },
        { dataType: 'enum', enums: ['confluence'] },
        { dataType: 'enum', enums: ['discord'] },
        { dataType: 'enum', enums: ['github'] },
        { dataType: 'enum', enums: ['hedera'] },
        { dataType: 'enum', enums: ['jira-software'] },
        { dataType: 'enum', enums: ['solana'] },
        { dataType: 'enum', enums: ['trello'] },
        { dataType: 'enum', enums: ['vercel'] },
      ],
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  StockCompany: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        priceChangedThreeMonthsPercent: { dataType: 'string', required: true },
        priceChangedOneMonthPercent: { dataType: 'string', required: true },
        priceChangedFiveDayPercent: { dataType: 'string', required: true },
        marketCap: { dataType: 'string', required: true },
        issueShare: { dataType: 'string', required: true },
        listedDate: { dataType: 'string', required: true },
        subsector: { dataType: 'string', required: true },
        sector: { dataType: 'string', required: true },
        supersector: { dataType: 'string', required: true },
        industry: { dataType: 'string', required: true },
        name: { dataType: 'string', required: true },
        market: { dataType: 'string', required: true },
        symbol: { dataType: 'string', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  StockHistory: {
    dataType: 'refAlias',
    type: {
      dataType: 'nestedObjectLiteral',
      nestedProperties: {
        timestamp: { dataType: 'string', required: true },
        volume: { dataType: 'string', required: true },
        close: { dataType: 'string', required: true },
        low: { dataType: 'string', required: true },
        high: { dataType: 'string', required: true },
        open: { dataType: 'string', required: true },
        symbol: { dataType: 'string', required: true },
        date: { dataType: 'string', required: true },
      },
      validators: {},
    },
  },
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
};
const validationService = new ValidationService(models);

// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

export function RegisterRoutes(app: express.Router) {
  // ###########################################################################################################
  //  NOTE: If you do not see routes for all of your controllers in this file, then you might not have informed tsoa of where to look
  //      Please look into the "controllerPathGlobs" config option described in the readme: https://github.com/lukeautry/tsoa
  // ###########################################################################################################
  app.get(
    '/api/crypto/coins',
    ...fetchMiddlewares<RequestHandler>(CryptoController),
    ...fetchMiddlewares<RequestHandler>(CryptoController.prototype.getCoins),

    function CryptoController_getCoins(request: any, response: any, next: any) {
      const args = {
        limit: { default: 100, in: 'query', name: 'limit', dataType: 'double' },
        offset: { default: 0, in: 'query', name: 'offset', dataType: 'double' },
        orderBy: {
          default: 'marketCap',
          in: 'query',
          name: 'orderBy',
          ref: 'OrderBy',
        },
        orderDirection: {
          default: 'desc',
          in: 'query',
          name: 'orderDirection',
          ref: 'OrderDirection',
        },
        tags: {
          default: '',
          in: 'query',
          name: 'tags',
          dataType: 'union',
          subSchemas: [{ ref: 'Tag' }, { dataType: 'string' }],
        },
        timePeriod: {
          default: '24h',
          in: 'query',
          name: 'timePeriod',
          ref: 'TimePeriod',
        },
        tier: { default: '1', in: 'query', name: 'tier', ref: 'Tier' },
      };

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new CryptoController();

        const promise = controller.getCoins.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, 200, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/forex/rates',
    ...fetchMiddlewares<RequestHandler>(ForexController),
    ...fetchMiddlewares<RequestHandler>(ForexController.prototype.getForex),

    function ForexController_getForex(request: any, response: any, next: any) {
      const args = {
        category: {
          default: null,
          in: 'query',
          name: 'category',
          ref: 'Category',
        },
      };

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new ForexController();

        const promise = controller.getForex.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, 200, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/health',
    ...fetchMiddlewares<RequestHandler>(HealthController),
    ...fetchMiddlewares<RequestHandler>(HealthController.prototype.get),

    function HealthController_get(request: any, response: any, next: any) {
      const args = {};

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new HealthController();

        const promise = controller.get.apply(controller, validatedArgs as any);
        promiseHandler(controller, promise, response, undefined, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/status',
    ...fetchMiddlewares<RequestHandler>(StatusesController),
    ...fetchMiddlewares<RequestHandler>(
      StatusesController.prototype.getStatuses
    ),

    function StatusesController_getStatuses(
      request: any,
      response: any,
      next: any
    ) {
      const args = {};

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new StatusesController();

        const promise = controller.getStatuses.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, 200, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/status/:service',
    ...fetchMiddlewares<RequestHandler>(StatusesController),
    ...fetchMiddlewares<RequestHandler>(
      StatusesController.prototype.getServiceStatus
    ),

    function StatusesController_getServiceStatus(
      request: any,
      response: any,
      next: any
    ) {
      const args = {
        service: {
          in: 'path',
          name: 'service',
          required: true,
          ref: 'Service',
        },
      };

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new StatusesController();

        const promise = controller.getServiceStatus.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, 200, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/vnindex/companies',
    ...fetchMiddlewares<RequestHandler>(VnindexController),
    ...fetchMiddlewares<RequestHandler>(
      VnindexController.prototype.getCompanies
    ),

    function VnindexController_getCompanies(
      request: any,
      response: any,
      next: any
    ) {
      const args = {};

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new VnindexController();

        const promise = controller.getCompanies.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, undefined, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/vnindex/history/:stockCode',
    ...fetchMiddlewares<RequestHandler>(VnindexController),
    ...fetchMiddlewares<RequestHandler>(VnindexController.prototype.getHistory),

    function VnindexController_getHistory(
      request: any,
      response: any,
      next: any
    ) {
      const args = {
        stockCode: {
          in: 'path',
          name: 'stockCode',
          required: true,
          dataType: 'string',
        },
        offset: { default: 0, in: 'query', name: 'offset', dataType: 'double' },
        limit: { default: 20, in: 'query', name: 'limit', dataType: 'double' },
      };

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new VnindexController();

        const promise = controller.getHistory.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, 200, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
  app.get(
    '/api/vnindex/history/:stockCode/chart',
    ...fetchMiddlewares<RequestHandler>(VnindexController),
    ...fetchMiddlewares<RequestHandler>(
      VnindexController.prototype.chartifyHistory
    ),

    function VnindexController_chartifyHistory(
      request: any,
      response: any,
      next: any
    ) {
      const args = {
        request: {
          in: 'request',
          name: 'request',
          required: true,
          dataType: 'object',
        },
        stockCode: {
          in: 'path',
          name: 'stockCode',
          required: true,
          dataType: 'string',
        },
        width: { default: 500, in: 'query', name: 'width', dataType: 'double' },
        height: {
          default: 100,
          in: 'query',
          name: 'height',
          dataType: 'double',
        },
        minValue: {
          default: 0,
          in: 'query',
          name: 'minValue',
          dataType: 'double',
        },
        maxValue: {
          default: 200,
          in: 'query',
          name: 'maxValue',
          dataType: 'double',
        },
        strokeColor: {
          default: '#0074d9',
          in: 'query',
          name: 'strokeColor',
          dataType: 'string',
        },
        strokeWidth: {
          default: 2,
          in: 'query',
          name: 'strokeWidth',
          dataType: 'double',
        },
      };

      // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

      let validatedArgs: any[] = [];
      try {
        validatedArgs = getValidatedArgs(args, request, response);

        const controller = new VnindexController();

        const promise = controller.chartifyHistory.apply(
          controller,
          validatedArgs as any
        );
        promiseHandler(controller, promise, response, 200, next);
      } catch (err) {
        return next(err);
      }
    }
  );
  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

  function isController(object: any): object is Controller {
    return (
      'getHeaders' in object && 'getStatus' in object && 'setStatus' in object
    );
  }

  function promiseHandler(
    controllerObj: any,
    promise: any,
    response: any,
    successStatus: any,
    next: any
  ) {
    return Promise.resolve(promise)
      .then((data: any) => {
        let statusCode = successStatus;
        let headers;
        if (isController(controllerObj)) {
          headers = controllerObj.getHeaders();
          statusCode = controllerObj.getStatus() || statusCode;
        }

        // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

        returnHandler(response, statusCode, data, headers);
      })
      .catch((error: any) => next(error));
  }

  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

  function returnHandler(
    response: any,
    statusCode?: number,
    data?: any,
    headers: any = {}
  ) {
    if (response.headersSent) {
      return;
    }
    Object.keys(headers).forEach((name: string) => {
      response.set(name, headers[name]);
    });
    if (
      data &&
      typeof data.pipe === 'function' &&
      data.readable &&
      typeof data._read === 'function'
    ) {
      data.pipe(response);
    } else if (data !== null && data !== undefined) {
      response.status(statusCode || 200).json(data);
    } else {
      response.status(statusCode || 204).end();
    }
  }

  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

  function responder(
    response: any
  ): TsoaResponse<HttpStatusCodeLiteral, unknown> {
    return function (status, data, headers) {
      returnHandler(response, status, data, headers);
    };
  }

  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa

  function getValidatedArgs(args: any, request: any, response: any): any[] {
    const fieldErrors: FieldErrors = {};
    const values = Object.keys(args).map((key) => {
      const name = args[key].name;
      switch (args[key].in) {
        case 'request':
          return request;
        case 'query':
          return validationService.ValidateParam(
            args[key],
            request.query[name],
            name,
            fieldErrors,
            undefined,
            { noImplicitAdditionalProperties: 'throw-on-extras' }
          );
        case 'path':
          return validationService.ValidateParam(
            args[key],
            request.params[name],
            name,
            fieldErrors,
            undefined,
            { noImplicitAdditionalProperties: 'throw-on-extras' }
          );
        case 'header':
          return validationService.ValidateParam(
            args[key],
            request.header(name),
            name,
            fieldErrors,
            undefined,
            { noImplicitAdditionalProperties: 'throw-on-extras' }
          );
        case 'body':
          return validationService.ValidateParam(
            args[key],
            request.body,
            name,
            fieldErrors,
            undefined,
            { noImplicitAdditionalProperties: 'throw-on-extras' }
          );
        case 'body-prop':
          return validationService.ValidateParam(
            args[key],
            request.body[name],
            name,
            fieldErrors,
            'body.',
            { noImplicitAdditionalProperties: 'throw-on-extras' }
          );
        case 'formData':
          if (args[key].dataType === 'file') {
            return validationService.ValidateParam(
              args[key],
              request.file,
              name,
              fieldErrors,
              undefined,
              { noImplicitAdditionalProperties: 'throw-on-extras' }
            );
          } else if (
            args[key].dataType === 'array' &&
            args[key].array.dataType === 'file'
          ) {
            return validationService.ValidateParam(
              args[key],
              request.files,
              name,
              fieldErrors,
              undefined,
              { noImplicitAdditionalProperties: 'throw-on-extras' }
            );
          } else {
            return validationService.ValidateParam(
              args[key],
              request.body[name],
              name,
              fieldErrors,
              undefined,
              { noImplicitAdditionalProperties: 'throw-on-extras' }
            );
          }
        case 'res':
          return responder(response);
      }
    });

    if (Object.keys(fieldErrors).length > 0) {
      throw new ValidateError(fieldErrors, '');
    }
    return values;
  }

  // WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
}

// WARNING: This file was auto-generated with tsoa. Please do not modify it. Re-run tsoa to re-generate this file: https://github.com/lukeautry/tsoa
