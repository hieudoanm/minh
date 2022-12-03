"""
Forex
"""


import requests


BASE_URL = 'https://raw.githubusercontent.com/hieudoanm/tablebase/master/json/forex'


CURRENCIES_URL = f"{BASE_URL}/currencies.json"


def get_symbols_currencies():
    """Get Symbols - Currencies"""
    response = requests.get(CURRENCIES_URL, timeout=10)
    return response.json()


def process_forex(words):
    """Process Forex"""
    symbols_currencies = get_symbols_currencies()
    symbols = list(
        map(lambda symbol_currency: symbol_currency["symbol"].lower(), symbols_currencies))
    filtered_symbols = list(set(words).intersection(symbols))
    data = {}
    if len(filtered_symbols) > 0:
        data["symbols"] = filtered_symbols
    return data
