import requests
import json
from . import utils


def test_error_no_jwt_token():
    url = "http://localhost:8080/api/v1/questions"
    headers = {
        "Content-Type": "application/json",
    }
    payload = json.dumps({"question": "what is promise?"})

    response = requests.request("POST", url, headers=headers, data=payload).json()

    assert response["data"] == None
    assert response["error"] != None
    assert "token" in response["error"]


def test_error_jwt_token_invalid():
    url = "http://localhost:8080/api/v1/questions"
    headers = {
        "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5MzIwNjksInVsZXJfaWQiOiIyInd.K1sEDl8Ke7VdYszxZd6ODH6DVZU7L4jTYs5GgeyIXIU",
        "Content-Type": "application/json",
    }
    payload = json.dumps({"question": "what is promise?"})

    response = requests.request("POST", url, headers=headers, data=payload).json()

    assert response["data"] == None
    assert response["error"] != None
    assert "token" in response["error"]
