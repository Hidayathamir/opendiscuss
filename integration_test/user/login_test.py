import requests
import json
from . import utils


headers = {"Content-Type": "application/json"}


def test_error_user_not_found():
    url = "http://localhost:8080/api/v1/login"
    payload = json.dumps(
        {"username": utils.generate_string(), "password": "dummy pass"}
    )

    response = requests.request("POST", url, headers=headers, data=payload).json()

    assert response["data"] == None
    assert "record not found" in response["error"]


def test_success():
    username = utils.generate_string()
    password = "dummy pass"
    utils.register_user(username, password)

    payload = json.dumps({"username": username, "password": password})

    url = "http://localhost:8080/api/v1/login"
    response = requests.request("POST", url, headers=headers, data=payload).json()

    assert response["data"] != None
    token: str = response["data"]["token"]
    assert token.startswith("Bearer")
    assert response["error"] == None
