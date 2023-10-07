from . import utils


headers = {"Content-Type": "application/json"}


def test_error_duplicate_username():
    username = utils.generate_string()
    password = "dummy pass"

    utils.register_user(username, password)
    response = utils.register_user(username, password)

    assert response["data"] == None
    assert "Duplicate entry" in response["error"]


def test_success():
    username = utils.generate_string()
    password = "dummy pass"

    response = utils.register_user(username, password)

    assert response["data"] != None
    assert response["data"]["user_id"] != 0
    assert response["error"] == None
