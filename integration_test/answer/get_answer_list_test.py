from integration_test import utils


def test_success():
    # same as `create_answer_test.py test_success`
    pass


def test_success_question_not_found():
    response = utils.get_answer_list_by_question_id(1000000000)

    assert response["data"] != None
    assert response["data"]["answers"] != None
    assert len(response["data"]["answers"]) == 0
    assert response["error"] == None


def test_should_sorted_by_thumbs_rate_desc():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")

    answer_id1 = utils.create_answer_get_id(jwt_token, question_id, "idk")
    answer_id2 = utils.create_answer_get_id(
        jwt_token, question_id, "software engineer is me"
    )
    utils.create_answer(jwt_token, question_id, "software engineer is us")

    utils.vote_answer(jwt_token, answer_id2, utils.VoteOption.ThumbsUp)
    utils.vote_answer(jwt_token, answer_id1, utils.VoteOption.ThumbsDown)

    response = utils.get_answer_list_by_question_id(question_id)

    assert response["data"] != None
    assert response["data"]["answers"] != None
    assert response["error"] == None

    answers = response["data"]["answers"]

    prev_thumbs_rate = 1000
    for _, answer in enumerate(answers):
        assert prev_thumbs_rate > answer["thumbs_rate"]
        prev_thumbs_rate = answer["thumbs_rate"]
