# opendiscuss

community platform

## user stories

- [x] user can register
- [x] user can login
- [x] user can ask question
- [x] user can see question, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count
- [x] user can vote thumbsup | thumbsdown question
- [x] user can edit (only) their question
- [x] user can delete (only) their question
- [x] user can answer question
- [x] user can see answer, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by highest thumbsup-thumbsdown count.
- [x] user can vote thumbsup | thumbsdown answer
- [x] user can edit (only) their answer
- [ ] user can delete (only) their answer
- [ ] user can comment on answer
- [ ] user can see comment, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by highest thumbsup-thumbsdown count.
- [ ] user can vote thumbsup | thumbsdown comment
- [ ] user can edit (only) their comment
- [ ] user can delete (only) their comment
- [ ] user can comment on comment (subcomment)
- [ ] user can see subcomment, with thumbsup-thumbsdown count, thumbsup count, thumbsdown count. Should be sorted by highest thumbsup-thumbsdown count.
- [ ] user can vote thumbsup | thumbsdown subcomment
- [ ] user can edit (only) their subcomment
- [ ] user can delete (only) their subcomment

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
