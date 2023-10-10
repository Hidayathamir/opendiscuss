from integration_test import utils


def test_error_question_not_found():
    response = utils.get_question_detail(1000000)

    assert response["data"] == None
    assert "record not found" in response["error"]


def test_success_create_question_then_get_question_detail():
    # same as `create_question_test.py test_success`
    pass
