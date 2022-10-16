import airVisualClient from '../libs/air-visual';

export const getAirVisualMessage = async (): Promise<string> => {
  const { data } = await airVisualClient.getAirQuality(
    'Vietnam',
    'Ho Chi Minh City',
    'Ho Chi Minh City'
  );
  return `Air Quality: ${data.current.pollution.aqius}`;
};
