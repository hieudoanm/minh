import { Article, SortBy } from '@hieudoanm/news';
import { addZero } from '@hieudoanm/utils';
import { newsClient } from '../libs/news';

const getMessage = (articles: Article[]) => {
  const border: string = [...Array(30).keys()].map(() => '-').join('');
  const list: string = articles
    .map((article: Article, index: number) => {
      const { source, title, url } = article;
      return `${addZero(index + 1)}. [${source.name} - ${title}](${url})`;
    })
    .join(`\n${border}\n`);
  return `${border}\n${list}\n${border}`;
};

export const getTopHeadlinesMessage = async (): Promise<string> => {
  const { articles = [] } = await newsClient.getTopHeadlines();
  return getMessage(articles);
};

export const getBlockchainCryptoNewsMessage = async (): Promise<string> => {
  const { articles = [] } = await newsClient.getEverything({
    sortBy: SortBy.POPULARITY,
    q: 'blockchain+crypto',
    pageSize: 10,
  });
  return getMessage(articles);
};
