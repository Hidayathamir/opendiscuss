from utils import utils


def test_success():
    # same as `create_answer_test.py test_success`
    pass


def test_error_answer_not_found():
    response = utils.get_answer_detail(1000000000)

    assert response["data"] == None
    assert response["error"] != None
    assert "record not found" in response["error"]
