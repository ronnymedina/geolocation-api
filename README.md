# Geolocation API :turtle:

Hi, if you are interested in this repository, I will tell you what it does.

### Enviroments

- GIN_MODE=(`debug` | `release`). Set `release` if you're in production.
- PORT
- POSTGRE_HOST
- POSTGRE_USER
- POSTGRE_PASS
- POSTGRE_DB
- POSTGRE_PORT

### Start app in dev

1) Start pg service: `docker-compose up -d geo-pg`
2) Create table: `docker exec -it geo-pg psql -U postgres -W your-database-here`. If the database does not exist, then you need to create it.
3) Copy the content of `migration.sql` in terminal.
4) Start api: `docker-compose up -d geo-api`

### Generate Documentation
If you have error with the command (`swag init -g`), see this [post](https://github.com/swaggo/swag/issues/197#issuecomment-585102353).

- Generate files: `swag init -g ./src/main.go -o ./src/docs`


### Main dependencies

- [Gin-gonic](https://github.com/gin-gonic/gin)
- [Field validator](https://pkg.go.dev/github.com/go-playground/validator/v10#readme-usage-and-documentation)
- [Gin swagger](https://github.com/swaggo/gin-swagger)
- [swag](https://github.com/swaggo/swag)

### Documentation
 - [postgresql-geo-queries-made-easy](https://postindustria.com/postgresql-geo-queries-made-easy/)
