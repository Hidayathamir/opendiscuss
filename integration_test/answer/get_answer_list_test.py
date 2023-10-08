from utils import utils


def test_success():
    # same as `create_answer_test.py test_success`
    pass


def test_success_question_not_found():
    response = utils.get_answer_list_by_question_id(1000000000)

    assert response["data"] != None
    assert response["data"]["answers"] != None
    assert len(response["data"]["answers"]) == 0
    assert response["error"] == None
