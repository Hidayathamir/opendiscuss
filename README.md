# opendiscuss

opendiscuss is community platform where people contribute to discuss some topic.

## user stories

You can see user story in file [README-user-story.md](https://github.com/Hidayathamir/opendiscuss/blob/master/README-user-story.md).

## entiry relationship diagram

You can see ERD in [dbdiagram](https://dbdiagram.io/d/opendiscuss-erd-64e854c002bd1c4a5e6392d5) and [dbdocs](https://dbdocs.io/hidayathamir/opendiscuss).

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
