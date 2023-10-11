# opendiscuss

community platform

## user stories

- [x] user can register
```shell
curl --location 'http://localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "hidayat",
    "password": "pass"
}'
```
- [x] user can login
```shell
curl --location 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "hidayat",
    "password": "pass"
}'
```
- [x] user can ask question
```shell
curl --location 'http://localhost:8080/api/v1/questions' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "question": "Mengapa Luhut diberi banyak jabatan oleh Presiden Joko Widodo?"
}'
```
- [x] user can see question, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by latest question first
```shell
curl --location 'http://localhost:8080/api/v1/questions'
```
```shell
curl --location 'http://localhost:8080/api/v1/questions/{{question_id}}'
```
- [x] user can vote thumbsup | thumbsdown question
```shell
curl --location --request POST 'http://localhost:8080/api/v1/questions/{{question_id}}/thumbsup' \
--header 'Authorization: Bearer {{jwt-token}}'
```
```shell
curl --location --request POST 'http://localhost:8080/api/v1/questions/{{question_id}}/thumbsdown' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can edit (only) their question
```shell
curl --location --request PUT 'http://localhost:8080/api/v1/questions/{{question_id}}' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "question": "update question"
}
'
```
- [x] user can delete (only) their question
```shell
curl --location --request DELETE 'http://localhost:8080/api/v1/questions/{{question_id}}' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can answer question
```shell
curl --location 'http://localhost:8080/api/v1/questions/{{question_id}}/answers' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "answer": "Doi selalu punya cara penyelesaiannya."
}'
```
- [x] user can see answer, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by highest thumbsup-thumbsdown count.
```shell
curl --location 'http://localhost:8080/api/v1/questions/{{question_id}}/answers'
```
```shell
curl --location 'http://localhost:8080/api/v1/answers/{{answer_id}}'
```
- [x] user can vote thumbsup | thumbsdown answer
```shell
curl --location --request POST 'http://localhost:8080/api/v1/answers/{{answer_id}}/thumbsup' \
--header 'Authorization: Bearer {{jwt-token}}'
```
```shell
curl --location --request POST 'http://localhost:8080/api/v1/answers/{{answer_id}}/thumbsdown' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can edit (only) their answer
```shell
curl --location --request PUT 'http://localhost:8080/api/v1/answers/{{answer_id}}' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "answer": "update answer"
}
'
```
- [x] user can delete (only) their answer
```shell
curl --location --request DELETE 'http://localhost:8080/api/v1/answers/{{answer_id}}' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can comment on answer
```shell
curl --location 'http://localhost:8080/api/v1/answers/{{answer_id}}/comments' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "comment": "Setuju gan."
}'
```
- [x] user can see comment, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by highest thumbsup-thumbsdown count.
```shell
curl --location 'http://localhost:8080/api/v1/answers/{{answer_id}}/comments'
```
```shell
curl --location 'http://localhost:8080/api/v1/comments/{{comment_id}}'
```
- [x] user can vote thumbsup | thumbsdown comment
```shell
curl --location --request POST 'http://localhost:8080/api/v1/comments/{{comment_id}}/thumbsup' \
--header 'Authorization: Bearer {{jwt-token}}'
```
```shell
curl --location --request POST 'http://localhost:8080/api/v1/comments/{{comment_id}}/thumbsdown' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can edit (only) their comment
```shell
curl --location --request PUT 'http://localhost:8080/api/v1/comments/{{comment_id}}' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "comment": "update comment"
}
'
```
- [x] user can delete (only) their comment
```shell
curl --location --request DELETE 'http://localhost:8080/api/v1/comments/{{comment_id}}/' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can comment on comment (subcomment)
```shell
curl --location 'http://localhost:8080/api/v1/answers/{{answer_id}}/comments' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "comment": "Lu mah setuju setuju aje.",
    "parent_comment_id": {{comment_id}}
}'
```
- [x] user can see subcomment, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by highest thumbsup-thumbsdown count.
```shell
curl --location 'http://localhost:8080/api/v1/comments/{{comment_id}}/subcomments'
```
```shell
curl --location 'http://localhost:8080/api/v1/comments/{{subcomment_id}}'
```
- [x] user can vote thumbsup | thumbsdown subcomment
```shell
curl --location --request POST 'http://localhost:8080/api/v1/comments/{{subcomment_id}}/thumbsup' \
--header 'Authorization: Bearer {{jwt-token}}'
```
```shell
curl --location --request POST 'http://localhost:8080/api/v1/comments/{{subcomment_id}}/thumbsdown' \
--header 'Authorization: Bearer {{jwt-token}}'
```
- [x] user can edit (only) their subcomment
```shell
curl --location --request PUT 'http://localhost:8080/api/v1/comments/{{subcomment_id}}' \
--header 'Authorization: Bearer {{jwt-token}}' \
--header 'Content-Type: application/json' \
--data '{
    "comment": "update comment"
}
'
```
- [x] user can delete (only) their subcomment
```shell
curl --location --request DELETE 'http://localhost:8080/api/v1/comments/{{subcomment_id}}/' \
--header 'Authorization: Bearer {{jwt-token}}'
```

## entiry relationship diagram

https://dbdiagram.io/d/opendiscuss-erd-64e854c002bd1c4a5e6392d5

https://dbdocs.io/hidayathamir/opendiscuss

## quick start

1. run mysql db using docker compose

```
sudo docker compose up
```

to see mysql credential you can check file `docker-compose.yml`

2. rename `.env-example` to `.env`

3. run golang

```
go run .
```

## api documentation

for api documentation, you can import file `postman/opendiscuss.postman_collection.json` into your postman.
