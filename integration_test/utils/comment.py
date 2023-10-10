import json
import requests
from integration_test.utils.vote_option import VoteOption


def create_comment(jwt_token, answer_id, comment):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}/comments"

    payload = json.dumps({"comment": comment})
    headers = {"Authorization": jwt_token, "Content-Type": "application/json"}

    return requests.request("POST", url, headers=headers, data=payload).json()


def create_comment_get_id(jwt_token, answer_id, comment):
    resposne = create_comment(jwt_token, answer_id, comment)
    return resposne["data"]["comment_id"]


def get_comment_detail(comment_id):
    url = f"http://localhost:8080/api/v1/comments/{comment_id}"

    return requests.request("GET", url, headers={}, data={}).json()


def get_comment_list_by_answer_id(answer_id):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}/comments"

    return requests.request("GET", url, headers={}, data={}).json()


def vote_comment(jwt_token, comment_id, voting: VoteOption):
    url = f"http://localhost:8080/api/v1/comments/{comment_id}/{voting.value}"

    headers = {"Authorization": jwt_token}

    return requests.request("POST", url, headers=headers, data={}).json()
