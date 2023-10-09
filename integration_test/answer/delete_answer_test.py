from utils import utils


def test_success():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy question")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "dummy answer")

    response = utils.delete_answer(jwt_token, answer_id)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] == None
    assert response["error"] != None
    assert "record not found" in response["error"]


def test_error_answer_not_found():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)

    response = utils.delete_answer(jwt_token, 10000000)

    assert response["data"] == None
    assert response["error"] != None
    assert "record not found" in response["error"]


def test_error_delete_other_user_answer():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(usera_jwt_token, "dummy question")
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, "dummy answer"
    )

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.delete_answer(userb_jwt_token, usera_answer_id)

    assert response["data"] == None
    assert response["error"] != None
    assert "unauthorized" in response["error"]
