package services

import (
	"ronnymedina/geolocation-api/src/config"
	"ronnymedina/geolocation-api/src/helpers"
	"ronnymedina/geolocation-api/src/validations"
)

func CreatePlace(place validations.CreatePlace) int64 {
	var id int64
	sql := "INSERT INTO places (name, description, resource_id, lat, lng) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	db := config.GetDB()
	err := db.QueryRow(sql, place.Name, place.Description, place.Resource_id, place.Lat, place.Lng).Scan(&id)

	helpers.CheckAndThrowException(err)

	return id
}
