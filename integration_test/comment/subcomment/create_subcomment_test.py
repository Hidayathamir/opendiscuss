from integration_test import utils


def test_success_create_subcomment_of_comment():
    username = utils.generate_string()
    password = utils.generate_string()
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is programmer"
    )
    comment_id = utils.create_comment_get_id(
        jwt_token, answer_id, "i do not agree with you"
    )

    subcomment = "why not agree with me?"
    response = utils.create_subcomment(jwt_token, answer_id, subcomment, comment_id)

    assert response["data"] != None
    assert response["data"]["comment_id"] != 0
    assert response["error"] == None

    subcomment_id = response["data"]["comment_id"]
    response = utils.get_comment_detail(subcomment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == subcomment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["author_id"] == user_id
    assert response["data"]["comment"]["comment"] == subcomment
    assert response["data"]["comment"]["thumbs_rate"] == 0
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["data"]["comment"]["comment_count"] == 0
    assert response["error"] == None

    response = utils.get_subcomment_list_by_comment_id(comment_id)

    assert response["data"] != None
    assert response["data"]["comments"] != None
    assert len(response["data"]["comments"]) == 1
    assert response["error"] == None


def test_error_create_subcomment_of_comment__comment_not_found():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is programmer"
    )

    response = utils.create_subcomment(
        jwt_token, answer_id, "why not agree with me?", 1000000
    )

    assert response["data"] == None
    assert response["error"] != None
    assert "a foreign key constraint fails" in response["error"]


def test_success_create_subcomment_of_subcomment():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is programmer"
    )
    comment_id = utils.create_comment_get_id(
        jwt_token, answer_id, "i do not agree with you"
    )
    subcomment_id = utils.create_subcomment_get_id(
        jwt_token, answer_id, "why not agree with me?", comment_id
    )

    sub_subcomment = "you are a liar"
    response = utils.create_subcomment(
        jwt_token, answer_id, sub_subcomment, subcomment_id
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] != 0
    assert response["error"] == None


def test_error_create_subcomment_of_subcomment__subcomment_not_found():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is programmer"
    )

    response = utils.create_subcomment(
        jwt_token, answer_id, "you are a liar", 100000000
    )

    assert response["data"] == None
    assert response["error"] != None
    assert "a foreign key constraint fails" in response["error"]
