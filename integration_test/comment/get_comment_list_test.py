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
