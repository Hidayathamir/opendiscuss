from utils import utils


def test_success():
    username = utils.generate_string()
    password = utils.generate_string()
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is SOLID?")
    answer = "SOLID is principle to achieve clean code"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    new_answer = "SOLID is scam"
    response = utils.update_answer(jwt_token, answer_id, new_answer)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["author_id"] == user_id
    assert response["data"]["answer"]["answer"] == new_answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 0


def test_error_update_other_user_question():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    usera_user_id = utils.register_user_get_id(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(usera_jwt_token, "what is SOLID?")
    usera_answer = "SOLID is principle to achieve clean code"
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, usera_answer
    )

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    userb_new_answer = "SOLID is scam"
    response = utils.update_answer(userb_jwt_token, usera_answer_id, userb_new_answer)

    assert response["data"] == None
    assert response["error"] != None
    assert "unauthorized" in response["error"]

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["author_id"] == usera_user_id
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 0
