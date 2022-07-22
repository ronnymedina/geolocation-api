# Geolocation API :turtle:

Hi, if you are interested in this repository, I will tell you what it does.

### Enviroments

- GIN_MODE=(`debug` | `release`). Set `release` if you're in production.
- PORT
- POSTGRE_HOST
- POSTGRE_USER
- POSTGRE_PASS
- POSTGRE_DB

### Commands

- start app: `go run ./src/main`
- Start app with docker-compose: `docker-compose up -d`


### Main dependencies

- [Gin-gonic](https://github.com/gin-gonic/gin)
- [Field validator](https://pkg.go.dev/github.com/go-playground/validator/v10#readme-usage-and-documentation)


### Documentation
 - [postgresql-geo-queries-made-easy](https://postindustria.com/postgresql-geo-queries-made-easy/)
