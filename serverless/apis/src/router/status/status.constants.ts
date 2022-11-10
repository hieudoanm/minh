import { Service } from './status.types';

export const SERVICES: Service[] = [
  'atlassian',
  'bitbucket',
  'circleci',
  'confluence',
  'discord',
  'github',
  'hedera',
  'jira-software',
  'solana',
  'trello',
  'vercel',
];

export const SERVICE_NAMES: Record<Service, string> = {
  atlassian: 'Atlassian',
  bitbucket: 'Bitbucket',
  circleci: 'Circle CI',
  confluence: 'Confluence',
  discord: 'Discord',
  github: 'GitHub',
  hedera: 'Hedera',
  'jira-software': 'Jira Software',
  solana: 'Solana',
  trello: 'Trello',
  vercel: 'Vercel',
};

export const SERVICE_URLS: Record<Service, string> = {
  atlassian: 'https://status.atlassian.com/api/v2/status.json',
  bitbucket: 'https://bitbucket.status.atlassian.com/api/v2/status.json',
  circleci: 'https://status.circleci.com/api/v2/status.json',
  confluence: 'https://confluence.status.atlassian.com/api/v2/status.json',
  discord: 'https://discordstatus.com/api/v2/status.json',
  github: 'https://www.githubstatus.com/api/v2/status.json',
  hedera: 'https://status.hedera.com/api/v2/status.json',
  'jira-software':
    'https://jira-software.status.atlassian.com/api/v2/status.json',
  solana: 'https://status.solana.com/api/v2/status.json',
  trello: 'https://trello.status.atlassian.com/api/v2/status.json',
  vercel: 'https://www.vercel-status.com/api/v2/status.json',
};
