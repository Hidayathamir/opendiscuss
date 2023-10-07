import requests


def test_success():
    url = "http://localhost:8080/api/v1/questions"

    response = requests.request("GET", url, headers={}, data={}).json()

    assert response["data"] != None
    assert response["data"]["questions"] != None
    assert response["error"] == None
