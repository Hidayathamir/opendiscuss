from integration_test import utils


def test_success():
    username = utils.generate_string()
    password = utils.generate_string()
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is SOLID?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "SOLID is principle to achieve clean code"
    )
    comment = "you are wrong"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    new_comment = "you are really wrong"
    response = utils.update_comment(jwt_token, comment_id, new_comment)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["author_id"] == user_id
    assert response["data"]["comment"]["comment"] == new_comment
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 0


def test_error_update_other_user_comment():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    usera_user_id = utils.register_user_get_id(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(usera_jwt_token, "what is SOLID?")
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, "SOLID is principle to achieve clean code"
    )
    usera_comment = "you are wrong"
    usera_comment_id = utils.create_comment_get_id(
        usera_jwt_token, usera_answer_id, usera_comment
    )

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    userb_new_comment = "you are really wrong"
    response = utils.update_comment(
        userb_jwt_token, usera_comment_id, userb_new_comment
    )

    assert response["data"] == None
    assert response["error"] != None
    assert "unauthorized" in response["error"]

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["author_id"] == usera_user_id
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 0
