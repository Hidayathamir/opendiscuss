from dateutil import parser
from integration_test import utils


def test_success():
    response = utils.get_question_list()

    assert response["data"] != None
    assert response["data"]["questions"] != None
    assert response["error"] == None


def test_should_sorted_by_created_at_desc():
    username = utils.generate_string()
    password = "dummy password"
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)

    utils.create_question_get_id(jwt_token, "dummy title", "first question")
    utils.create_question_get_id(jwt_token, "dummy title", "second question")

    response = utils.get_question_list()

    assert response["data"] != None
    assert response["data"]["questions"] != None
    assert response["error"] == None

    questions = response["data"]["questions"]

    prev_created_at = parser.parse("2030-10-10T14:54:55Z")
    for _, question in enumerate(questions):
        created_at = parser.parse(question["created_at"])
        assert prev_created_at >= created_at
        prev_created_at = created_at
