import json
import requests
from integration_test.utils.vote_option import VoteOption


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


def update_answer(jwt_token, answer_id, answer):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}"

    payload = json.dumps({"answer": answer})
    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }

    return requests.request("PUT", url, headers=headers, data=payload).json()


def delete_answer(jwt_token, answer_id):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}"

    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }

    return requests.request("DELETE", url, headers=headers, data={}).json()
