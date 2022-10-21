import { Match, Status } from '@hieudoanm/football';
import logger from '@hieudoanm/pino';
import { footballClient } from '../libs/football';

const TIME_ZONE = 'Asia/Ho_Chi_Minh';

const getLastMatchMessage = ({ matches }: { matches: Match[] }): string => {
  const finishedMatches: Match[] = matches.filter(
    (match: Match) => match.status === Status.FINISHED
  );
  const lastMatch: Match = finishedMatches[finishedMatches.length - 1];
  logger.info('Last Match', { lastMatch });
  const {
    utcDate = '',
    score: {
      fullTime: { home: homeScore = '0', away: awayScore = '0' },
    },
    homeTeam: { name: homeName = '' },
    awayTeam: { name: awayName = '' },
  } = lastMatch;
  const d = new Date(utcDate);
  const dateTime: string = d.toLocaleString('en', { timeZone: TIME_ZONE });
  return `${dateTime}\n${homeName} ${homeScore} - ${awayScore} ${awayName}`;
};

const getNextMatchMessage = ({ matches }: { matches: Match[] }): string => {
  const scheduledMatches = matches.filter(
    (match: Match) => match.status === 'TIMED'
  );
  const nextMatch: Match = scheduledMatches[0];
  logger.info('Next Match', { nextMatch });
  const {
    utcDate = '',
    homeTeam: { name: homeName = '' },
    awayTeam: { name: awayName = '' },
  } = nextMatch;
  const d = new Date(utcDate);
  const dateTime: string = d.toLocaleString('en', { timeZone: TIME_ZONE });
  return `${dateTime}\n${homeName} - ${awayName}`;
};

export const getMatchesMessage = async () => {
  const { matches = [] } = await footballClient.getMatchesByTeam(64);
  const lastMatchMessage = getLastMatchMessage({ matches });
  const nextMatchMessage = getNextMatchMessage({ matches });
  const border: string = [...Array(30).keys()].map(() => '-').join('');
  return `${lastMatchMessage}\n${border}\n${nextMatchMessage}`;
};
