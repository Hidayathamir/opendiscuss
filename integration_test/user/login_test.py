from integration_test import utils


headers = {"Content-Type": "application/json"}


def test_error_user_not_found():
    username = utils.generate_string()
    password = "dummy pass"

    response = utils.login_user(username, password)

    assert response["data"] == None
    assert "record not found" in response["error"]


def test_success():
    username = utils.generate_string()
    password = "dummy pass"
    utils.register_user(username, password)

    response = utils.login_user(username, password)

    assert response["data"] != None
    token: str = response["data"]["token"]
    assert token.startswith("Bearer")
    assert response["error"] == None
