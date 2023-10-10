from integration_test import utils


def test_vote_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "who are you?")
    answer = "i am hidayat"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "who are you?")
    answer = "i am hidayat"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 1
    assert response["error"] == None


def test_vote_thumbs_up_then_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "who are you?")
    answer = "i am hidayat"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_up_then_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "who are you?")
    answer = "i am hidayat"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 1
    assert response["error"] == None


def test_vote_thumbs_down_then_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "who are you?")
    answer = "i am hidayat"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 1
    assert response["error"] == None

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_down_then_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "who are you?")
    answer = "i am hidayat"
    answer_id = utils.create_answer_get_id(jwt_token, question_id, answer)

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 1
    assert response["error"] == None

    response = utils.vote_answer(jwt_token, answer_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["answer_id"] == answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == answer_id
    assert response["data"]["answer"]["author"] == username
    assert response["data"]["answer"]["answer"] == answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None


def test_usera_thumbs_up_userb_thumbs_up():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(usera_jwt_token, "who are you?")
    usera_answer = "i am hidayat"
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, usera_answer
    )

    response = utils.vote_answer(
        usera_jwt_token, usera_answer_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["answer_id"] == usera_answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_answer(
        userb_jwt_token, usera_answer_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["answer_id"] == usera_answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 2
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None


def test_usera_thumbs_up_userb_thumbs_down():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(usera_jwt_token, "who are you?")
    usera_answer = "i am hidayat"
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, usera_answer
    )

    response = utils.vote_answer(
        usera_jwt_token, usera_answer_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["answer_id"] == usera_answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 0
    assert response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_answer(
        userb_jwt_token, usera_answer_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["answer_id"] == usera_answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 1
    assert response["data"]["answer"]["thumbs_down"] == 1
    assert response["error"] == None


def test_usera_thumbs_down_userb_thumbs_down():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(usera_jwt_token, "who are you?")
    usera_answer = "i am hidayat"
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, usera_answer
    )

    response = utils.vote_answer(
        usera_jwt_token, usera_answer_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["answer_id"] == usera_answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 1
    assert response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_answer(
        userb_jwt_token, usera_answer_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["answer_id"] == usera_answer_id
    assert response["error"] == None

    response = utils.get_answer_detail(usera_answer_id)

    assert response["data"] != None
    assert response["data"]["answer"]["id"] == usera_answer_id
    assert response["data"]["answer"]["author"] == usera_username
    assert response["data"]["answer"]["answer"] == usera_answer
    assert response["data"]["answer"]["thumbs_up"] == 0
    assert response["data"]["answer"]["thumbs_down"] == 2
    assert response["error"] == None
