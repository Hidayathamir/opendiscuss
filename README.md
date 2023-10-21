# opendiscuss

Opendiscuss is a community platform where people contribute to discuss various topics.

## User Stories

You can find user stories in the file [README-user-story.md](README-user-story.md).

## Entiry Relationship Diagram

You can view the ERD on [dbdiagram](https://dbdiagram.io/d/opendiscuss-erd-64e854c002bd1c4a5e6392d5) and [dbdocs](https://dbdocs.io/hidayathamir/opendiscuss).

![erd](README_asset/erd.png)

## Quick Start

You can check this youtube video for [Quick Start Opendiscuss](https://youtu.be/S9iGM_GFVkc?si=zawTc9gu3J-Yw99J).

1. Rename `.env-example` to `.env`.

2. Run the MySQL database and the Go app using Docker Compose:

```shell
sudo docker compose up --build
```

MySQL credentials is:

```
DB_USER="user"
DB_PASSWORD="password"
DB_HOST="localhost"
DB_PORT="9306"
DB_NAME="opendiscuss"
```

Test hit API login:

```shell
curl --location 'http://localhost:9080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "abdan",
    "password": "abdan"
}'
```

Check [README-user-story.md](README-user-story.md) for other API.

## API Documentation

For API documentation, you can import the file [postman/opendiscuss.postman_collection.json](postman/opendiscuss.postman_collection.json) into your Postman.
