from integration_test import utils


def test_success():
    username = utils.generate_string()
    password = utils.generate_string()
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is programmer"
    )

    comment = "i do not agree with you, you are wrong."
    response = utils.create_comment(jwt_token, answer_id, comment)

    assert response["data"] != None
    assert response["data"]["comment_id"] != 0
    assert response["error"] == None

    comment_id = response["data"]["comment_id"]
    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["author_id"] == user_id
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 0
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["data"]["comment"]["comment_count"] == 0
    assert response["error"] == None

    response = utils.get_comment_list_by_answer_id(answer_id)

    assert response["data"] != None
    assert response["data"]["comments"] != None
    assert len(response["data"]["comments"]) == 1
    assert response["error"] == None


def test_error_answer_not_found():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)

    comment = "software engineer is programmer"
    response = utils.create_comment(jwt_token, 1000000000, comment)

    assert response["data"] == None
    assert response["error"] != None
    assert "foreign key constraint fails" in response["error"]
