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


def login_user(username, password):
    url = "http://localhost:8080/api/v1/login"
    headers = {"Content-Type": "application/json"}
    payload = json.dumps({"username": username, "password": password})

    response = requests.request("POST", url, headers=headers, data=payload).json()

    return response["data"]["token"]


def create_question(jwt_token, question):
    url = "http://localhost:8080/api/v1/questions"
    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }
    payload = json.dumps({"question": question})

    response = requests.request("POST", url, headers=headers, data=payload).json()

    return response["data"]["question_id"]


def get_question_detail(questoin_id):
    url = f"http://localhost:8080/api/v1/questions/{questoin_id}"

    return requests.request("GET", url, headers={}, data={}).json()
