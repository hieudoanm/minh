"""
Weather
"""


import requests
import unidecode
from nltk import pos_tag

from .nlp import process_ngrams


def process_weather(words):
    """Process Weather"""
    capitialized_words = list(map(lambda word: word.capitalize(), words))
    capitialized_sentence = " ".join(capitialized_words)
    ngrams_words = process_ngrams(capitialized_sentence, 4)
    tagged_words = pos_tag(ngrams_words)
    nnp_words = list(map(
        lambda tagged_word: tagged_word[0],
        list(
            filter(lambda tagged_word: tagged_word[1] == "NNP", tagged_words))
    ))
    cities = process_cities(nnp_words)
    data = {}
    data["cities"] = cities
    return data


CITIES_URL = "https://raw.githubusercontent.com/hieudoanm/tablebase/master/json/world/cities.json"


def is_city(word, name) -> bool:
    """Is City"""
    number_of_words = len(word.split(" "))
    if number_of_words == 1 or number_of_words == 2:
        return word.lower() == name.lower()
    elif number_of_words >= 3:
        return word.lower() in name.lower()
    return False


def process_cities(words):
    """Process Cities"""
    response = requests.get(CITIES_URL, timeout=10)
    response_json = response.json()
    names = list(map(lambda city: city.get("name"), response_json))
    unidecoded_names = list(map(unidecode.unidecode, names))
    unique_unidecoded_names = list(set(unidecoded_names))
    unique_unidecoded_names.sort()
    cities = []
    for word in words:
        results = []
        for name in unique_unidecoded_names:
            if is_city(word, name):
                results.append(name)
        cities += results
    return sorted(cities, key=len, reverse=True)
