from . import utils


def test_question_not_found():
    response = utils.get_question_detail(1000000)

    assert response["data"] == None
    assert "record not found" in response["error"]


def test_success():
    response = utils.get_question_detail(1)

    assert response["data"] != None
    question_id: int = response["data"]["question"]["id"]
    assert question_id == 1
    assert response["error"] == None


def test_success_create_question_then_get_question_detail():
    username = utils.generate_string()
    password = "dummy password"
    utils.register_user(username, password)
    jwt_token = utils.login_user(username, password)
    question = "what is stackoverflow?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["error"] == None
