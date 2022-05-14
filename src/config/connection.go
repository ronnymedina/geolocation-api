package config

import (
	"database/sql"
	"os"
	"ronnymedina/geolocation-api/src/helpers"
)

var DB *sql.DB
var PG_HOST = os.Getenv("POSTGRE_HOST")
var PG_USER = os.Getenv("POSTGRE_USER")
var PG_PASS = os.Getenv("POSTGRE_PASS")
var PG_DB = os.Getenv("POSTGRE_DB")

func GetDB() *sql.DB {
	return DB
}

func StartConnection() {
	url := "postgres://" + PG_USER + ":" + PG_PASS + "@" + PG_HOST + "/" + PG_DB + "?sslmode=disable"
	db, err := sql.Open("postgres", url)

	helpers.CheckAndThrowException(err)

	err = db.Ping()
	helpers.CheckAndThrowException(err)

	DB = db
}
