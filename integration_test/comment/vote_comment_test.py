from integration_test import utils


def test_vote_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy title", "who are you?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "i am hidayat")
    comment = "you are AI"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 1
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy title", "who are you?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "i am hidayat")
    comment = "you are AI"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == -1
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 1
    assert response["error"] == None


def test_vote_thumbs_up_then_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy title", "who are you?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "i am hidayat")
    comment = "you are AI"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 1
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 0
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_up_then_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy title", "who are you?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "i am hidayat")
    comment = "you are AI"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 1
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == -1
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 1
    assert response["error"] == None


def test_vote_thumbs_down_then_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy title", "who are you?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "i am hidayat")
    comment = "you are AI"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == -1
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 1
    assert response["error"] == None

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 1
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_down_then_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "dummy title", "who are you?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "i am hidayat")
    comment = "you are AI"
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, comment)

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == -1
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 1
    assert response["error"] == None

    response = utils.vote_comment(jwt_token, comment_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["comment_id"] == comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == comment_id
    assert response["data"]["comment"]["author"] == username
    assert response["data"]["comment"]["comment"] == comment
    assert response["data"]["comment"]["thumbs_rate"] == 0
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None


def test_usera_thumbs_up_userb_thumbs_up():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(
        usera_jwt_token, "dummy title", "who are you?"
    )
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, "i am hidayat"
    )
    usera_comment = "you are AI"
    usera_comment_id = utils.create_comment_get_id(
        usera_jwt_token, usera_answer_id, usera_comment
    )

    response = utils.vote_comment(
        usera_jwt_token, usera_comment_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] == usera_comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_rate"] == 1
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_comment(
        userb_jwt_token, usera_comment_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] == usera_comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_rate"] == 2
    assert response["data"]["comment"]["thumbs_up"] == 2
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None


def test_usera_thumbs_up_userb_thumbs_down():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(
        usera_jwt_token, "dummy title", "who are you?"
    )
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, "i am hidayat"
    )
    usera_comment = "you are AI"
    usera_comment_id = utils.create_comment_get_id(
        usera_jwt_token, usera_answer_id, usera_comment
    )

    response = utils.vote_comment(
        usera_jwt_token, usera_comment_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] == usera_comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_rate"] == 1
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 0
    assert response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_comment(
        userb_jwt_token, usera_comment_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] == usera_comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_rate"] == 0
    assert response["data"]["comment"]["thumbs_up"] == 1
    assert response["data"]["comment"]["thumbs_down"] == 1
    assert response["error"] == None


def test_usera_thumbs_down_userb_thumbs_down():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question_id = utils.create_question_get_id(
        usera_jwt_token, "dummy title", "who are you?"
    )
    usera_answer_id = utils.create_answer_get_id(
        usera_jwt_token, usera_question_id, "i am hidayat"
    )
    usera_comment = "you are AI"
    usera_comment_id = utils.create_comment_get_id(
        usera_jwt_token, usera_answer_id, usera_comment
    )

    response = utils.vote_comment(
        usera_jwt_token, usera_comment_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] == usera_comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_rate"] == -1
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 1
    assert response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_comment(
        userb_jwt_token, usera_comment_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["comment_id"] == usera_comment_id
    assert response["error"] == None

    response = utils.get_comment_detail(usera_comment_id)

    assert response["data"] != None
    assert response["data"]["comment"]["id"] == usera_comment_id
    assert response["data"]["comment"]["author"] == usera_username
    assert response["data"]["comment"]["comment"] == usera_comment
    assert response["data"]["comment"]["thumbs_rate"] == -2
    assert response["data"]["comment"]["thumbs_up"] == 0
    assert response["data"]["comment"]["thumbs_down"] == 2
    assert response["error"] == None
