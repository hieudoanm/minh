export type Service =
  | 'atlassian'
  | 'bitbucket'
  | 'circleci'
  | 'confluence'
  | 'discord'
  | 'github'
  | 'hedera'
  | 'jira-software'
  | 'solana'
  | 'trello'
  | 'vercel';

export type Status = 'active' | 'inactive';

export type ServiceStatusResponse = { status: { indicator: string } };

export type StatusResponse = { service: string; status: Status };

export type StatusesResponse = Record<Service, StatusResponse>;
