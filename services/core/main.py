"""
Natural Language Processing as a Service
"""

import os
import re

import nltk
from dotenv import load_dotenv
from flask import Flask, Response, jsonify, request
# Libs
from libs.football import process_football
from libs.forex import process_forex
from libs.news import process_news
from libs.weather import process_weather
# NLTK
from nltk.corpus import stopwords
from nltk.tokenize import word_tokenize

nltk.download("averaged_perceptron_tagger")
nltk.download("punkt")
nltk.download("stopwords")

stop_words = set(stopwords.words("english"))

app = Flask(__name__)


@ app.route("/", methods=["GET", "POST"])
def home() -> Response:
    """Home Page"""
    status = "healthy"
    response = {}
    response["status"] = status
    return jsonify(response)


@ app.route("/health", methods=["GET", "POST"])
def health() -> Response:
    """Health Check"""
    status = "healthy"
    response = {}
    response["status"] = status
    return jsonify(response)


@ app.route("/categorise", methods=["POST"])
def categorise() -> Response:
    """Categorise Word"""
    request_body = request.json
    return jsonify(request_body)


@ app.route("/summarise", methods=["POST"])
def summarise() -> Response:
    """Summarise Paragraph"""
    request_body = request.json
    return jsonify(request_body)


@ app.route("/process", methods=["POST"])
def process() -> Response:
    """Process Message"""
    request_body = request.json
    message = request_body.get("message", "")
    lower_message = message.lower()
    no_symbol_string = re.sub(r"[^\w\s]", "", lower_message)
    words = word_tokenize(no_symbol_string)
    filtered_words = [w for w in words if w not in stop_words]
    intent = get_intent(filtered_words)
    data = process_intent(intent, words)
    response_body = {
        "intent": intent,
        "data": data,
    }
    return jsonify(response_body)


def get_intent(words) -> str:
    """Get Intent"""
    all_intents = ["football", "forex", "news", "weather"]
    all_intents.sort()
    intents = list(set(all_intents).intersection(words))
    return intents[0]


def process_intent(intent, words) -> dict:
    """Process Intent"""
    if intent == "football":
        return process_football(words)
    elif intent == "forex":
        return process_forex(words)
    elif intent == "news":
        return process_news(words)
    if intent == "weather":
        return process_weather(words)
    return {}


if __name__ == "__main__":
    load_dotenv()
    ENVIRONMENT = os.getenv("ENVIRONMENT").lower()
    PORT = os.getenv("PORT")
    if ENVIRONMENT == "development":
        app.run(port=PORT, debug=True)
    else:
        from waitress import serve
        serve(app, port=PORT)
