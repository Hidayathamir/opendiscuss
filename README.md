# opendiscuss

## user stories

[opendiscuss mvp stories](https://docs.google.com/document/d/1QaN062BwKNMS8vxxDgYFGQ3SQkWDC67etQO1ZacL6d4/edit?usp=sharing).

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
