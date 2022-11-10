/* eslint-disable no-console */
/* eslint-disable @typescript-eslint/no-var-requires */
const { readFileSync, writeFileSync } = require('fs');
const path = require('path');
const { convert } = require('swagger2-to-postmanv2');

const convertSync = (swaggerJson) => {
  return new Promise((resolve, reject) => {
    convert(swaggerJson, { collectionName: 'test' }, (error, result) => {
      if (error) {
        return reject(error);
      }
      resolve(result);
    });
  });
};

const saveOutput = (filePath, output) => {
  const { type, data } = output;
  if (type === 'collection') {
    const collectionFilePath = path.join(__dirname, filePath);
    writeFileSync(collectionFilePath, JSON.stringify(data, null, 2));
  }
};

const main = async () => {
  try {
    console.log(__dirname);
    const swaggerFilePath = path.join(__dirname, '../../docs/swagger.json');
    console.log('swagger file path', swaggerFilePath);
    const swaggerString = readFileSync(swaggerFilePath, { encoding: 'utf-8' });
    const swaggerJson = JSON.parse(swaggerString);
    console.log('swagger.json', swaggerJson);
    const postmanResult = await convertSync({
      type: 'json',
      data: { ...swaggerJson, swagger: '2.0' },
    });
    console.log('postman result', postmanResult);
    const { result, output = [] } = postmanResult;
    if (!result) return;
    // If output is only one
    if (output.length === 1) {
      saveOutput(`../../postman/collections/main.json`, output[0]);
    } else {
      // If there is more than 1 output
      for (let i = 0; i < output.length; i++) {
        saveOutput(`../../postman/collections/main_${i}.json`, output[i]);
      }
    }

    process.exit(0);
  } catch (error) {
    console.error(error);
    process.exit(1);
  }
};

main().catch((error) => console.error('Error', error));
