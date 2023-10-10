import json
import requests


def create_comment(jwt_token, answer_id, comment):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}/comments"

    payload = json.dumps({"comment": comment})
    headers = {"Authorization": jwt_token, "Content-Type": "application/json"}

    return requests.request("POST", url, headers=headers, data=payload).json()


def create_comment_get_id(jwt_token, question_id, answer):
    resposne = create_comment(jwt_token, question_id, answer)
    return resposne["data"]["comment_id"]


def get_comment_detail(comment_id):
    url = f"http://localhost:8080/api/v1/comments/{comment_id}"

    return requests.request("GET", url, headers={}, data={}).json()


def get_comment_list_by_answer_id(answer_id):
    url = f"http://localhost:8080/api/v1/answers/{answer_id}/comments"

    return requests.request("GET", url, headers={}, data={}).json()
