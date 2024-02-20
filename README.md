# okidoks-server

`okidoks-server` is meant to provide a simple API for exchanging data with
its sister project [okidoks-client](https://github.com/mugraph/okidoks-client).

1. Have a postgres database named `okidoks_db` and user `postgres` on
   `localhost` and bound to port `5432`.
2. Run `go run main.go` and head over to `localhost:8080/publications`.

## Prerequisites

- [gin](https://github.com/gin-gonic/gin): router & webserver
- [gorm](https://gorm.io/gorm): object-relational mapper (ORM)
- [postgres](https://www.postgresql.org/): database

## REST-API

### GET Methods

- `/resources` returns all resources
- `/resource/prefix/*suffix` returns single resource
- `/contributors` returns all contributors
- `/publishers` returns all publishers

### POST Methods

- `/resource` creates a new resource with `{ URL: string }`
