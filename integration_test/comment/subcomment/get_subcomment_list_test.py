from integration_test import utils


def test_get_subcomment_list():
    # covered in create_subcomment_test.py
    pass


def test_should_sorted_by_thumbs_rate_desc():
    username = utils.generate_string()
    password = utils.generate_string()
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question_id = utils.create_question_get_id(jwt_token, "what is software engineer?")
    answer_id = utils.create_answer_get_id(jwt_token, question_id, "idk")
    comment_id = utils.create_comment_get_id(jwt_token, answer_id, "hmm")

    subcom1 = utils.create_subcomment_get_id(
        jwt_token, answer_id, "subcom1", comment_id
    )
    subcom2 = utils.create_subcomment_get_id(
        jwt_token, answer_id, "subcom2", comment_id
    )
    utils.create_subcomment(jwt_token, answer_id, "subcom3", comment_id)

    utils.vote_comment(jwt_token, subcom2, utils.VoteOption.ThumbsUp)
    utils.vote_comment(jwt_token, subcom1, utils.VoteOption.ThumbsDown)

    response = utils.get_subcomment_list_by_comment_id(comment_id)

    assert response["data"] != None
    assert response["data"]["comments"] != None
    assert response["error"] == None

    comments = response["data"]["comments"]

    prev_thumbs_rate = 1000
    for _, comment in enumerate(comments):
        assert prev_thumbs_rate > comment["thumbs_rate"]
        prev_thumbs_rate = comment["thumbs_rate"]
