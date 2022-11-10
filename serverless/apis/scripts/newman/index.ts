import logger from '@hieudoanm/pino';
import newman, { NewmanRunSummary } from 'newman';
import { Collection, CollectionDefinition } from 'postman-collection';
import postmanTestCollection from '../../postman/collections/test.json';
import localEnvVar from '../../postman/environments/local.json';

type EnvVar = { key: string; value: string };

const TEST_ENVIRONMENT = process.env.TEST_ENVIRONMENT || 'local';

const envVarOptions: Record<string, EnvVar[]> = {
  local: localEnvVar,
};

const getFileName = (): string => {
  const [date, time] = new Date().toISOString().split('T');
  const [hours, minutes] = time.split(':');
  return `${date}-${hours}-${minutes}`;
};

const runSync = async (
  collection: Collection | CollectionDefinition | string
): Promise<NewmanRunSummary> => {
  const envVar = envVarOptions[TEST_ENVIRONMENT];
  const timestamp = getFileName();
  return new Promise((resolve, reject) => {
    newman.run(
      {
        envVar,
        collection,
        reporters: ['cli', 'html'],
        reporter: {
          html: { export: `./coverage/newman-report-${timestamp}.html` },
        },
      },
      (error: Error | null, summary: NewmanRunSummary) => {
        if (error) {
          reject(error);
        }
        resolve(summary);
      }
    );
  });
};

const main = async () => {
  await runSync(postmanTestCollection);
  process.exit(0);
};

main().catch((error) => logger.error('Error', { error }));
