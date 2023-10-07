from . import utils


def test_success():
    username = utils.generate_string()
    password = "dummy password"
    utils.register_user(username, password)
    jwt_token = utils.login_user_get_token(username, password)
    question = "what is promise?"

    response = utils.create_question(jwt_token, question)

    assert response["data"] != None
    assert response["data"]["question_id"] != 0
    assert response["error"] == None
