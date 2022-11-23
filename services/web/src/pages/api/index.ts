import type { NextApiRequest, NextApiResponse } from 'next';

type StatusData = {
  status: string;
};

const handler = (
  _request: NextApiRequest,
  response: NextApiResponse<StatusData>
) => {
  response.status(200).json({ status: 'healthy' });
};

export default handler;
