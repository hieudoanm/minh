"""
News
"""


from nltk import pos_tag

from .nlp import process_ngrams

NEWS_CATEGORIES = ['businessent', 'ertainment', 'general',
                   'health', 'science', 'sports', 'technology']


NEWS_SOURCES = {
    'abc': {'id': 'abc-news', 'name': 'ABC News'},
    'abc news': {'id': 'abc-news', 'name': 'ABC News'},
    'abc news au': {'id': 'abc-news-au', 'name': 'ABC News (AU)'},
    'aftenposten': {'id': 'aftenposten', 'name': 'Aftenposten'},
    'al jazeera english': {'id': 'al-jazeera-english', 'name': 'Al Jazeera English'},
    'ansa': {'id': 'ansa', 'name': 'ANSA.it'},
    'argaam': {'id': 'argaam', 'name': 'Argaam'},
    'ars technica': {'id': 'ars-technica', 'name': 'Ars Technica'},
    'ary news': {'id': 'ary-news', 'name': 'Ary News'},
    'associated press': {'id': 'associated-press', 'name': 'Associated Press'},
    'australian financial review': {'id': 'australian-financial-review',
                                    'name': 'Australian Financial Review'},
    'axios': {'id': 'axios', 'name': 'Axios'},
    'bbc': {'id': 'bbc-news', 'name': 'BBC News'},
    'bbc news': {'id': 'bbc-news', 'name': 'BBC News'},
    'bbc sport': {'id': 'bbc-sport', 'name': 'BBC Sport'},
    'bild': {'id': 'bild', 'name': 'Bild'},
    'blasting': {'id': 'blasting-news-br', 'name': 'Blasting News (BR)'},
    'blasting news': {'id': 'blasting-news-br', 'name': 'Blasting News (BR)'},
    'bleacher report': {'id': 'bleacher-report', 'name': 'Bleacher Report'},
    'bloomberg': {'id': 'bloomberg', 'name': 'Bloomberg'},
    'breitbart': {'id': 'breitbart-news', 'name': 'Breitbart News'},
    'breitbart news': {'id': 'breitbart-news', 'name': 'Breitbart News'},
    'business insider': {'id': 'business-insider', 'name':  'Business Insider'},
    'business insider uk': {'id': 'business-insider-uk', 'name': 'Business Insider (UK)'},
    'buzzfeed': {'id': 'buzzfeed', 'name':  'Buzzfeed'},
    'cbc': {'id': 'cbc-news', 'name': 'CBC News'},
    'cbc news': {'id': 'cbc-news', 'name':  'CBC News'},
    'cbs': {'id': 'cbs-news', 'name': 'CBS News'},
    'cbs news': {'id': 'cbs-news', 'name': 'CBS News'},
    'cnn': {'id': 'cnn', 'name': 'CNN'},
    'cnn spanish': {'id': 'cnn-es', 'name': 'CNN Spanish'},
    'crypto coins': {'id': 'crypto-coins-news', 'name': 'Crypto Coins News'},
    'crypto coins news': {'id': 'crypto-coins-news', 'name': 'Crypto Coins News'},
    'der tagesspiegel': {'id': 'der-tagesspiegel', 'name': 'Der Tagesspiegel'},
    'die zeit': {'id': 'die-zeit', 'name': 'Die Zeit'},
    'el mundo': {'id': 'el-mundo', 'name': 'El Mundo'},
    'engadget': {'id': 'engadget', 'name': 'Engadget'},
    'entertainment weekly': {'id': 'entertainment-weekly', 'name': 'Entertainment Weekly'},
    'espn': {'id': 'espn', 'name': 'ESPN'},
    'espn cric info': {'id': 'espn-cric-info', 'name': 'ESPN Cric Info'},
    'financial post': {'id': 'financial-post', 'name': 'Financial Post'},
    'focus': {'id': 'focus', 'name': 'Focus'},
    'football italia': {'id': 'football-italia', 'name': 'Football Italia'},
    'fortune': {'id': 'fortune', 'name': 'Fortune'},
    '442': {'id': 'four-four-two', 'name': 'FourFourTwo'},
    'four four two': {'id': 'four-four-two', 'name': 'FourFourTwo'},
    'fox': {'id': 'fox-news', 'name': 'Fox News'},
    'fox news': {'id': 'fox-news', 'name': 'Fox News'},
    'fox sports': {'id': 'fox-sports', 'name': 'Fox Sports'},
    'globo': {'id': 'globo', 'name': 'Globo'},
    'google news': {'id': 'google-news', 'name': 'Google News'},
    'google news argentina': {'id': 'google-news-ar', 'name': 'Google News (Argentina)'},
    'google news australia': {'id': 'google-news-au', 'name': 'Google News (Australia)'},
    'google news brasil': {'id': 'google-news-br', 'name': 'Google News (Brasil)'},
    'google news canada': {'id': 'google-news-ca', 'name': 'Google News (Canada)'},
    'google news france': {'id': 'google-news-fr', 'name': 'Google News (France)'},
    'google news india': {'id': 'google-news-in', 'name': 'Google News (India)'},
    'google news israel': {'id': 'google-news-is', 'name': 'Google News (Israel)'},
    'google news italy': {'id': 'google-news-it', 'name': 'Google News (Italy)'},
    'google news russia': {'id': 'google-news-ru', 'name': 'Google News (Russia)'},
    'google news saudi arabia': {'id': 'google-news-sa', 'name': 'Google News (Saudi Arabia)'},
    'google news uk': {'id': 'google-news-uk', 'name': 'Google News (UK)'},
    'göteborgs posten': {'id': 'goteborgs-posten', 'name': 'Göteborgs-Posten'},
    'gruenderszene': {'id': 'gruenderszene', 'name': 'Gruenderszene'},
    'hacker news': {'id': 'hacker-news', 'name': 'Hacker News'},
    'handelsblatt': {'id': 'handelsblatt', 'name': 'Handelsblatt'},
    'ign': {'id': 'ign', 'name': 'IGN'},
    'il sole 24 ore': {'id': 'il-sole-24-ore', 'name': 'Il Sole 24 Ore'},
    'independent': {'id': 'independent', 'name': 'Independent'},
    'infobae': {'id': 'infobae', 'name': 'Infobae'},
    'infomoney': {'id': 'info-money', 'name': 'InfoMoney'},
    'la gaceta': {'id': 'la-gaceta', 'name': 'La Gaceta'},
    'la nacion': {'id': 'la-nacion', 'name': 'La Nacion'},
    'la repubblica': {'id': 'la-repubblica', 'name': 'La Repubblica'},
    'le monde': {'id': 'le-monde', 'name': 'Le Monde'},
    'lenta': {'id': 'lenta', 'name': 'Lenta'},
    "l'equipe": {'id': 'lequipe', 'name': "L'equipe"},
    'les echos': {'id': 'les-echos', 'name': 'Les Echos'},
    'libération': {'id': 'liberation', 'name': 'Libération'},
    'marca': {'id': 'marca', 'name': 'Marca'},
    'mashable': {'id': 'mashable', 'name': 'Mashable'},
    'medical news today': {'id': 'medical-news-today', 'name': 'Medical News Today'},
    'msnbc': {'id': 'msnbc', 'name': 'MSNBC'},
    'mtv': {'id': 'mtv-news', 'name': 'MTV News'},
    'mtv news': {'id': 'mtv-news', 'name': 'MTV News'},
    'mtv news uk': {'id': 'mtv-news-uk', 'name': 'MTV News (UK)'},
    'national geographic': {'id': 'national-geographic', 'name': 'National Geographic'},
    'national review': {'id': 'national-review', 'name': 'National Review'},
    'nbc news': {'id': 'nbc-news', 'name': 'NBC News'},
    'news24': {'id': 'news24', 'name': 'News24'},
    'news scientist': {'id': 'new-scientist', 'name': 'New Scientist'},
    'news.com.au': {'id': 'news-com-au', 'name': 'News.com.au'},
    'newsweek': {'id': 'newsweek', 'name': 'Newsweek'},
    'new york magazine': {'id': 'new-york-magazine', 'name': 'New York Magazine'},
    'next big future': {'id': 'next-big-future', 'name': 'Next Big Future'},
    'nfl': {'id': 'nfl-news', 'name': 'NFL News'},
    'nfl news': {'id': 'nfl-news', 'name': 'NFL News'},
    'nhl': {'id': 'nhl-news', 'name': 'NHL News'},
    'nhl news': {'id': 'nhl-news', 'name': 'NHL News'},
    'nrk': {'id': 'nrk', 'name': 'NRK'},
    'politico': {'id': 'politico', 'name': 'Politico'},
    'polygon': {'id': 'polygon', 'name': 'Polygon'},
    'rbc': {'id': 'rbc', 'name': 'RBC'},
    'recode': {'id': 'recode', 'name': 'Recode'},
    'reddit': {'id': 'reddit-r-all', 'name': 'Reddit / r/all'},
    'reuters': {'id': 'reuters', 'name': 'Reuters'},
    'rt': {'id': 'rt', 'name': 'RT'},
    'rte': {'id': 'rte', 'name': 'RTE'},
    'rtl nieuws': {'id': 'rtl-nieuws', 'name': 'RTL Nieuws'},
    'sabq': {'id': 'sabq', 'name': 'SABQ'},
    'spiegel online': {'id': 'spiegel-online', 'name': 'Spiegel Online'},
    'svenska dagbladet': {'id': 'svenska-dagbladet', 'name': 'Svenska Dagbladet'},
    't3n': {'id': 't3n', 'name': 'T3n'},
    'talk sport': {'id': 'talksport', 'name': 'TalkSport'},
    'tech crunch': {'id': 'techcrunch', 'name': 'TechCrunch'},
    'tech crunch cn': {'id': 'techcrunch-cn', 'name': 'TechCrunch (CN)'},
    'tech radar': {'id': 'techradar', 'name': 'TechRadar'},
    'the american conservative': {'id': 'the-american-conservative',
                                  'name': 'The American Conservative'},
    'the globe and mail': {'id': 'the-globe-and-mail', 'name': 'The Globe And Mail'},
    'the hill': {'id': 'the-hill', 'name': 'The Hill'},
    'the hindu': {'id': 'the-hindu', 'name': 'The Hindu'},
    'the huffington post': {'id': 'the-huffington-post', 'name': 'The Huffington Post'},
    'the irish times': {'id': 'the-irish-times', 'name': 'The Irish Times'},
    'the jerusalem post': {'id': 'the-jerusalem-post', 'name': 'The Jerusalem Post'},
    'the lad bible': {'id': 'the-lad-bible', 'name': 'The Lad Bible'},
    'the next web': {'id': 'the-next-web', 'name': 'The Next Web'},
    'the sport bible': {'id': 'the-sport-bible', 'name': 'The Sport Bible'},
    'the times of india': {'id': 'the-times-of-india', 'name': 'The Times of India'},
    'the verge': {'id': 'the-verge', 'name': 'The Verge'},
    'the wall street journal': {'id': 'the-wall-street-journal', 'name': 'The Wall Street Journal'},
    'the washington post': {'id': 'the-washington-post', 'name': 'The Washington Post'},
    'the washington times': {'id': 'the-washington-times', 'name': 'The Washington Times'},
    'time': {'id': 'time', 'name': 'Time'},
    'usa today': {'id': 'usa-today', 'name': 'USA Today'},
    'vice': {'id': 'vice-news', 'name': 'Vice News'},
    'vice news': {'id': 'vice-news', 'name': 'Vice News'},
    'wired': {'id': 'wired', 'name': 'Wired'},
    'wired de': {'id': 'wired-de', 'name': 'Wired.de'},
    'wirtschafts woche': {'id': 'wirtschafts-woche', 'name': 'Wirtschafts Woche'},
    'xinhua net': {'id': 'xinhua-net', 'name': 'Xinhua Net'},
    'ynet': {'id': 'ynet', 'name': 'Ynet'}
}


def process_news(words):
    """Process News"""
    data = {}
    # Process Categories
    categories = list(set(words).intersection(NEWS_CATEGORIES))
    if len(categories) > 0:
        data["categories"] = categories
    # Process Sources
    all_news_sources_keys = NEWS_SOURCES.keys()
    sentence = " ".join(words)
    ngrams_words = process_ngrams(sentence, 4)
    news_sources_keys = list(
        set(ngrams_words).intersection(all_news_sources_keys))
    if len(news_sources_keys) > 0:
        data["sources"] = list(
            map(lambda key: NEWS_SOURCES[key]['id'], news_sources_keys))
    # Process Query
    tagged_words = pos_tag(words)
    nn_words = list(map(
        lambda tagged_word: tagged_word[0],
        list(filter(lambda tagged_word: tagged_word[1] == "NN", tagged_words))
    ))
    nn_words = list(set(nn_words) - set(news_sources_keys) - set(['news']))
    if len(nn_words) > 0:
        data["query"] = nn_words
    return data
