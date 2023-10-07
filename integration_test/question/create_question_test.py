import requests
import json
from . import utils


def test_success():
    username = utils.generate_string()
    password = "dummy password"
    utils.register_user(username, password)
    jwt_token = utils.login_user(username, password)

    url = "http://localhost:8080/api/v1/questions"
    headers = {
        "Authorization": jwt_token,
        "Content-Type": "application/json",
    }
    payload = json.dumps({"question": "what is promise?"})

    response = requests.request("POST", url, headers=headers, data=payload).json()

    assert response["data"] != None
    assert response["data"]["question_id"] != 0
    assert response["error"] == None
