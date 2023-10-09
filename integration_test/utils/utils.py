from enum import Enum
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


def register_user_get_id(username, password):
    response = register_user(username, password)
    return response["data"]["user_id"]


def login_user(username, password):
    url = "http://localhost:8080/api/v1/login"
    headers = {"Content-Type": "application/json"}
    payload = json.dumps({"username": username, "password": password})
    return requests.request("POST", url, headers=headers, data=payload).json()


def login_user_get_token(username, password):
    response = login_user(username, password)
    return response["data"]["token"]


def create_question(jwt_token, question):
    url = "http://localhost:8080/api/v1/questions"
    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }
    payload = json.dumps({"question": question})

    return requests.request("POST", url, headers=headers, data=payload).json()


def create_question_get_id(jwt_token, question):
    response = create_question(jwt_token, question)
    return response["data"]["question_id"]


def get_question_detail(question_id):
    url = f"http://localhost:8080/api/v1/questions/{question_id}"

    return requests.request("GET", url, headers={}, data={}).json()


class VoteOption(Enum):
    ThumbsUp = "thumbsup"
    ThumbsDown = "thumbsdown"


def vote_question(jwt_token, question_id, voting: VoteOption):
    url = f"http://localhost:8080/api/v1/questions/{question_id}/{voting.value}"

    headers = {"Authorization": jwt_token}

    return requests.request("POST", url, headers=headers, data={}).json()


def update_question(jwt_token, question_id, question):
    url = f"http://localhost:8080/api/v1/questions/{question_id}"

    payload = json.dumps({"question": question})
    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }

    return requests.request("PUT", url, headers=headers, data=payload).json()


def delete_question(jwt_token, question_id):
    url = f"http://localhost:8080/api/v1/questions/{question_id}"

    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }

    return requests.request("DELETE", url, headers=headers, data={}).json()


def create_answer(jwt_token, question_id, answer):
    url = f"http://localhost:8080/api/v1/questions/{question_id}/answers"

    payload = json.dumps({"answer": answer})
    headers = {"Authorization": jwt_token, "Content-Type": "application/json"}

    return requests.request("POST", url, headers=headers, data=payload).json()


def create_answer_get_id(jwt_token, question_id, answer):
    resposne = create_answer(jwt_token, question_id, answer)
    return resposne["data"]["answer_id"]


def get_answer_detail(answer_id):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}"

    return requests.request("GET", url, headers={}, data={}).json()


def get_answer_list_by_question_id(question_id):
    url = f"http://localhost:8080/api/v1/questions/{question_id}/answers"

    return requests.request("GET", url, headers={}, data={}).json()


def vote_answer(jwt_token, answer_id, voting: VoteOption):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}/{voting.value}"

    headers = {"Authorization": jwt_token}

    return requests.request("POST", url, headers=headers, data={}).json()
