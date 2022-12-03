"""
Natural Language Processing
"""


from nltk import ngrams


def process_ngrams(sentence, number):
    """Process ngrams"""
    words = []
    for i in range(1, number + 1):
        n_grams = ngrams(sentence.split(), i)
        n_words = list(map(" ".join, n_grams))
        words += n_words
    return words
