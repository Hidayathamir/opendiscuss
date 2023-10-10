from integration_test import utils


def test_success():
    # same as `create_comment_test.py test_success`
    pass


def test_success_question_not_found():
    response = utils.get_comment_list_by_answer_id(1000000000)

    assert response["data"] != None
    assert response["data"]["comments"] != None
    assert len(response["data"]["comments"]) == 0
    assert response["error"] == None


def test_should_only_get_root_comment_of_answer():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is programmer"
    )
    comment_id = utils.create_comment_get_id(
        jwt_token, answer_id, "i do not agree with you, you are wrong."
    )
    utils.create_subcomment(jwt_token, answer_id, "hmmm....", comment_id)

    response = utils.get_comment_list_by_answer_id(answer_id)

    assert response["data"] != None
    assert response["data"]["comments"] != None
    assert len(response["data"]["comments"]) == 1
    assert response["error"] == None


def test_should_sorted_by_thumbs_rate_desc():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "idk")

    comment_id1 = utils.create_comment_get_id(jwt_token, answer_id, "hmm")
    comment_id2 = utils.create_comment_get_id(jwt_token, answer_id, "hmmmmm2")
    utils.create_comment(jwt_token, answer_id, "hmmmmmmmmm3")

    utils.vote_comment(jwt_token, comment_id2, utils.VoteOption.ThumbsUp)
    utils.vote_comment(jwt_token, comment_id1, utils.VoteOption.ThumbsDown)

    response = utils.get_comment_list_by_answer_id(answer_id)

    assert response["data"] != None
    assert response["data"]["comments"] != None
    assert response["error"] == None

    comments = response["data"]["comments"]

    prev_thumbs_rate = 1000
    for _, comment in enumerate(comments):
        assert prev_thumbs_rate > comment["thumbs_rate"]
        prev_thumbs_rate = comment["thumbs_rate"]
