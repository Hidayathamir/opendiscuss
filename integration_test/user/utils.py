import json
from random import choice
import requests
from string import ascii_uppercase


def generate_string(len=12):
    return "".join(choice(ascii_uppercase) for _ in range(len))


def register_user(username, password):
    url = "http://localhost:8080/api/v1/register"
    headers = {"Content-Type": "application/json"}
    payload = json.dumps({"username": username, "password": password})
    return requests.request("POST", url, headers=headers, data=payload).json()
