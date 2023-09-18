# okidoks-server
`okidoks-server` is meant to provide a simple API for exchanging data with
its sister project [okidoks-client](https://github.com/mugraph/okidoks-client).

## Prerequisites

- [gin](https://github.com/gin-gonic/gin): router & webserver
- [gorm](https://gorm.io/gorm): object-relational mapper (ORM)
- [postgres](https://www.postgresql.org/): database

## REST-API

### GET Methods
- `/publications` returns all available publications 

