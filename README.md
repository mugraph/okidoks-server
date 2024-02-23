# okidoks-server

`okidoks-server` is meant to provide a simple API for exchanging data with
its sister project [okidoks-client](https://github.com/mugraph/okidoks-client).

1. Have a postgres database named `okidoks_db` and user `postgres` on
   `localhost` and bound to port `5432`.
2. Build and run `./main` and head over to `localhost:8081/resources`.

## Prerequisites

- [gin](https://github.com/gin-gonic/gin): router & webserver
- [gorm](https://gorm.io/gorm): object-relational mapper (ORM)
- [postgres](https://www.postgresql.org/): database

## REST-API

| Method  | Endpoint                        | Description               |
|---------|---------------------------------|---------------------------|
| GET     |`/api/v1/resources`              | returns all resources     |
| GET     |`/api/v1/resource/prefix/*suffix`| returns single resource   |
| GET     |`/api/v1/contributors`           | returns all contributors  |
| GET     |`/api/v1/publishers`             | returns all publishers    |
| POST    |`/api/v1/resource`               | creates a new resource    |
