from utils import utils


def test_error_question_not_found():
    response = utils.get_question_detail(1000000)

    assert response["data"] == None
    assert "record not found" in response["error"]


def test_success_create_question_then_get_question_detail():
    username = utils.generate_string()
    password = "dummy password"
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "what is stackoverflow?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["author_id"] == user_id
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None
