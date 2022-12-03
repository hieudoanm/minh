"""
Football
"""


import requests

from .nlp import process_ngrams

BASE_URL = 'https://raw.githubusercontent.com/hieudoanm/tablebase/master/json/football'


COMPETITIONS_URL = f"{BASE_URL}/competitions.json"


def get_competitions():
    """Get Competitions"""
    response = requests.get(COMPETITIONS_URL, timeout=10)
    return response.json()


TEAMS_URL = f"{BASE_URL}/teams.json"


def get_teams():
    """Get Teams"""
    response = requests.get(TEAMS_URL, timeout=10)
    return response.json()


def process_football(words):
    """Process Football"""
    data = {}
    # Process Competitions
    competitions = process_competitions(words)
    if len(competitions) > 0:
        data["competitions"] = [
            item for sublist in competitions for item in sublist]
    # Process Teams
    teams = process_teams(words)
    if len(teams) > 0:
        data["teams"] = [item for sublist in teams for item in sublist]
    # Return Data
    return data


def process_competitions(words):
    """Process Competitions"""
    competitions = get_competitions()
    names = list(map(lambda team: team.get("name").lower(), competitions))
    ngrams_words = process_ngrams(" ".join(words), 4)
    competition_names = []
    for word in ngrams_words:
        results = []
        for name in names:
            if is_team(word, name):
                results.append(name)
        competition_names += results
    competition_names = list(set(competition_names))
    if len(competition_names) > 0:
        return list(
            map(lambda name: process_team(
                name, competitions), competition_names)
        )
    return []


def is_competition(word, name) -> bool:
    """Is Team"""
    number_of_words = len(word.split(" "))
    if number_of_words == 1 or number_of_words == 2:
        return word.lower() == name.lower()
    elif number_of_words >= 3:
        return word.lower() in name.lower()
    return False


def process_competition(name, competitions):
    """Process Team"""
    filtered_competitions = list(
        filter(lambda team: team["name"].lower() == name, competitions))
    competitions = []
    for competition in filtered_competitions:
        competitions.append({
            "id": int(competition["id"]),
            "name": competition["name"]
        })
    return competitions


def process_teams(words):
    """Process Teams"""
    teams = get_teams()
    names = list(map(lambda team: team.get("name").lower(), teams))
    short_names = list(map(lambda team: team.get("shortName").lower(), teams))
    all_names = names + short_names
    ngrams_words = process_ngrams(" ".join(words), 4)
    short_team_names = []
    for word in ngrams_words:
        results = []
        for name in all_names:
            if is_team(word, name):
                results.append(name)
        short_team_names += results
    short_team_names = list(set(short_team_names))
    if len(short_team_names) > 0:
        return list(
            map(lambda short_team_name: process_team(
                short_team_name, teams), short_team_names)
        )
    return []


def is_team(word, name) -> bool:
    """Is Team"""
    number_of_words = len(word.split(" "))
    if number_of_words == 1 or number_of_words == 2:
        return word.lower() == name.lower()
    elif number_of_words >= 3:
        return word.lower() in name.lower()
    return False


def process_team(short_team_name, teams):
    """Process Team"""
    filtered_teams = list(
        filter(lambda team:
               team["name"].lower() == short_team_name or
               team["shortName"].lower() == short_team_name, teams))
    teams = []
    for team in filtered_teams:
        teams.append({
            "id": int(team["id"]),
            "name": team["name"],
            "shortName": team["shortName"]
        })
    return teams
