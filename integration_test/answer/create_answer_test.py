from utils import utils


def test_success():
    username = utils.generate_string()
    password = utils.generate_string()
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "what is software engineer?"
    question_id = utils.create_question_get_id(jwt_token, question)

    answer = "software engineer is programmer"
    response = utils.create_answer(jwt_token, question_id, answer)

    assert response["data"] != None
    assert response["data"]["answer_id"] != 0
    assert response["error"] == None

    answer_id = response["data"]["answer_id"]
    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["author_id"] == user_id
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.get_answer_list_by_question_id(question_id)

    assert response["data"] != None
    assert response["data"]["answers"] != None
    assert len(response["data"]["answers"]) == 1
    assert response["error"] == None


def test_error_question_not_found():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)

    answer = "software engineer is programmer"
    response = utils.create_answer(jwt_token, 1000000000, answer)

    assert response["data"] == None
    assert response["error"] != None
    assert "foreign key constraint fails" in response["error"]
