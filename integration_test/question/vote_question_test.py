from utils import utils


def test_vote_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "who are you?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 1
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "who are you?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 1
    assert response["error"] == None


def test_vote_thumbs_up_then_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "who are you?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 1
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_up_then_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "who are you?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 1
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 1
    assert response["error"] == None


def test_vote_thumbs_down_then_thumbs_up():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "who are you?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 1
    assert response["error"] == None

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsUp)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 1
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None


def test_vote_thumbs_down_then_thumbs_down():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "who are you?"
    question_id = utils.create_question_get_id(jwt_token, question)

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 1
    assert response["error"] == None

    response = utils.vote_question(jwt_token, question_id, utils.VoteOption.ThumbsDown)

    assert response["data"] != None
    assert response["data"]["question_id"] == question_id
    assert response["error"] == None

    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == question_id
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None


def test_usera_thumbs_up_userb_thumbs_up():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question = "who are you?"
    usera_question_id = utils.create_question_get_id(usera_jwt_token, usera_question)

    usera_response = utils.vote_question(
        usera_jwt_token, usera_question_id, utils.VoteOption.ThumbsUp
    )

    assert usera_response["data"] != None
    assert usera_response["data"]["question_id"] == usera_question_id
    assert usera_response["error"] == None

    usera_response = utils.get_question_detail(usera_question_id)

    assert usera_response["data"] != None
    assert usera_response["data"]["question"]["id"] == usera_question_id
    assert usera_response["data"]["question"]["author"] == usera_username
    assert usera_response["data"]["question"]["question"] == usera_question
    assert usera_response["data"]["question"]["thumbs_up"] == 1
    assert usera_response["data"]["question"]["thumbs_down"] == 0
    assert usera_response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_question(
        userb_jwt_token, usera_question_id, utils.VoteOption.ThumbsUp
    )

    assert response["data"] != None
    assert response["data"]["question_id"] == usera_question_id
    assert response["error"] == None

    response = utils.get_question_detail(usera_question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == usera_question_id
    assert response["data"]["question"]["author"] == usera_username
    assert response["data"]["question"]["question"] == usera_question
    assert response["data"]["question"]["thumbs_up"] == 2
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None


def test_usera_thumbs_up_userb_thumbs_down():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question = "who are you?"
    usera_question_id = utils.create_question_get_id(usera_jwt_token, usera_question)

    usera_response = utils.vote_question(
        usera_jwt_token, usera_question_id, utils.VoteOption.ThumbsUp
    )

    assert usera_response["data"] != None
    assert usera_response["data"]["question_id"] == usera_question_id
    assert usera_response["error"] == None

    usera_response = utils.get_question_detail(usera_question_id)

    assert usera_response["data"] != None
    assert usera_response["data"]["question"]["id"] == usera_question_id
    assert usera_response["data"]["question"]["author"] == usera_username
    assert usera_response["data"]["question"]["question"] == usera_question
    assert usera_response["data"]["question"]["thumbs_up"] == 1
    assert usera_response["data"]["question"]["thumbs_down"] == 0
    assert usera_response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_question(
        userb_jwt_token, usera_question_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["question_id"] == usera_question_id
    assert response["error"] == None

    response = utils.get_question_detail(usera_question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == usera_question_id
    assert response["data"]["question"]["author"] == usera_username
    assert response["data"]["question"]["question"] == usera_question
    assert response["data"]["question"]["thumbs_up"] == 1
    assert response["data"]["question"]["thumbs_down"] == 1
    assert response["error"] == None


def test_usera_thumbs_down_userb_thumbs_down():
    usera_username = utils.generate_string()
    usera_password = utils.generate_string()
    utils.register_user(usera_username, usera_password)
    usera_jwt_token = utils.login_user_get_token(usera_username, usera_password)
    usera_question = "who are you?"
    usera_question_id = utils.create_question_get_id(usera_jwt_token, usera_question)

    usera_response = utils.vote_question(
        usera_jwt_token, usera_question_id, utils.VoteOption.ThumbsDown
    )

    assert usera_response["data"] != None
    assert usera_response["data"]["question_id"] == usera_question_id
    assert usera_response["error"] == None

    usera_response = utils.get_question_detail(usera_question_id)

    assert usera_response["data"] != None
    assert usera_response["data"]["question"]["id"] == usera_question_id
    assert usera_response["data"]["question"]["author"] == usera_username
    assert usera_response["data"]["question"]["question"] == usera_question
    assert usera_response["data"]["question"]["thumbs_up"] == 0
    assert usera_response["data"]["question"]["thumbs_down"] == 1
    assert usera_response["error"] == None

    userb_username = utils.generate_string()
    userb_password = utils.generate_string()
    utils.register_user(userb_username, userb_password)
    userb_jwt_token = utils.login_user_get_token(userb_username, userb_password)

    response = utils.vote_question(
        userb_jwt_token, usera_question_id, utils.VoteOption.ThumbsDown
    )

    assert response["data"] != None
    assert response["data"]["question_id"] == usera_question_id
    assert response["error"] == None

    response = utils.get_question_detail(usera_question_id)

    assert response["data"] != None
    assert response["data"]["question"]["id"] == usera_question_id
    assert response["data"]["question"]["author"] == usera_username
    assert response["data"]["question"]["question"] == usera_question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 2
    assert response["error"] == None
