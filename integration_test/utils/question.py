import json
import requests
from integration_test.utils.vote_option import VoteOption


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
