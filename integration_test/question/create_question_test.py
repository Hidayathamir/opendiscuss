from utils import utils


def test_success():
    username = utils.generate_string()
    password = "dummy password"
    user_id = utils.register_user_get_id(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "what is promise?"

    response = utils.create_question(jwt_token, question)

    assert response["data"] != None
    assert response["data"]["question_id"] != 0
    assert response["error"] == None

    question_id = response["data"]["question_id"]
    response = utils.get_question_detail(question_id)

    assert response["data"] != None
    assert response["data"]["question"]["author"] == username
    assert response["data"]["question"]["author_id"] == user_id
    assert response["data"]["question"]["question"] == question
    assert response["data"]["question"]["thumbs_up"] == 0
    assert response["data"]["question"]["thumbs_down"] == 0
    assert response["error"] == None
