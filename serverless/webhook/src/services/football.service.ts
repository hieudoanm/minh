import { Match } from '@hieudoanm/football';
import { footballClient } from '../libs/football';

export const getNextMatchMessage = async () => {
  const { matches = [] } = await footballClient.getMatchesByTeam(64);
  const nextMatch: {
    utcDate: string;
    status: string;
    homeTeam: { name: string };
    awayTeam: { name: string };
  } = matches.filter((match: Match) => match.status === 'TIMED')[0];
  console.log('Next Match', nextMatch);
  const {
    utcDate = '',
    homeTeam: { name: homeName = '' },
    awayTeam: { name: awayName = '' },
  } = nextMatch;
  const d = new Date(utcDate);
  const timeZone = 'Asia/Ho_Chi_Minh';
  const dateTime: string = d.toLocaleString('en', { timeZone });
  return `${dateTime}\n${homeName} - ${awayName}`;
};
