from integration_test import utils


def test_success():
    username = utils.generate_string()
    password = utils.generate_string()
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    title = "dummy title"
    question = "what is SOLID?"
    question_id = utils.create_question_get_id(jwt_token, title, question)

    new_title = "dummy title2"
    new_question = "who are you?"
    response = utils.update_question(jwt_token, question_id, new_title, new_question)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["author_id"] == user_id
    assert response["data"]["question"]["title"] == new_title
    assert response["data"]["question"]["question"] == new_question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 0


def test_error_update_other_user_question():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    usera_user_id = utils.register_user_get_id(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_title = "dummy title"
    usera_question = "what is SOLID?"
    usera_question_id = utils.create_question_get_id(
        usera_jwt_token, usera_title, usera_question
    )

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    userb_new_title = "dummy title2"
    userb_new_question = "who are you?"
    response = utils.update_question(
        userb_jwt_token, usera_question_id, userb_new_title, userb_new_question
    )

    assert response["data"] == None
    assert response["error"] != None
    assert "unauthorized" in response["error"]

    response = utils.get_question_detail(usera_question_id)

    assert response["data"]["question"]["id"] == usera_question_id
    assert response["data"]["question"]["author"] == usera_username
    assert response["data"]["question"]["author_id"] == usera_user_id
    assert response["data"]["question"]["title"] == usera_title
    assert response["data"]["question"]["question"] == usera_question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 0
