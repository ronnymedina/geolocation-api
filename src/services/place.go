package services

import (
	"log"
	"ronnymedina/geolocation-api/src/config"
	"ronnymedina/geolocation-api/src/helpers"
	"ronnymedina/geolocation-api/src/models"
	"ronnymedina/geolocation-api/src/validations"
	"time"
)

const (
	DATE_FORMAT = "2006-01-02 15:04:05"
)

func FindPlace(id int64) *models.Place {
	var place models.Place
	sql := "SELECT id, name, description, resource_id, lat, lng, created_at FROM places WHERE id = $1 AND is_deleted = false LIMIT 1"
	db := config.GetDB()
	err := db.QueryRow(sql, id).Scan(&place.Id, &place.Name, &place.Description, &place.ResourceId, &place.Lat, &place.Lng, &place.CreatedAt)

	helpers.CheckAndThrowException(err)

	return &place
}

func CreatePlace(place validations.CreatePlace) int64 {
	var id int64
	sql := "INSERT INTO places (name, description, resource_id, lat, lng) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	db := config.GetDB()
	err := db.QueryRow(sql, place.Name, place.Description, place.ResourceId, place.Lat, place.Lng).Scan(&id)

	helpers.CheckAndThrowException(err)

	return id
}

func UpdatePlace(id int64, place validations.UpdatePlace) {
	sql := "UPDATE places SET name = $1, description = $2, resource_id = $3, updated_at = $4 WHERE id = $5"
	db := config.GetDB()
	stmt, err := db.Prepare(sql)

	helpers.CheckAndThrowException(err)
	defer stmt.Close()

	_, err = stmt.Exec(place.Name, place.Description, place.ResourceId, time.Now().UTC().Format(DATE_FORMAT), id)
	helpers.CheckAndThrowException(err)
}

func DeletePlace(id int64) {
	sql := "UPDATE places SET is_deleted = $1 WHERE id = $2"
	stmt, err := config.GetDB().Prepare(sql)

	helpers.CheckAndThrowException(err)
	defer stmt.Close()

	_, err = stmt.Exec(true, id)
	helpers.CheckAndThrowException(err)
}

func GetNearbyPlaces(params validations.NearbyPlaces) []*validations.NearbyPlaceResult {
	// $1 = lat | $2 = lng | $3 = range | $4 = limit | $5 = page
	sql := `
	SELECT
		p.id AS id,
		p.name AS name,
		p.description AS description,
		p.resource_id AS resource_id,
		p.lat AS lat,
		p.lng AS lng,
		(ROUND(earth_distance(ll_to_earth($1, $2), ll_to_earth(p.lat, p.lng))::NUMERIC, 2)/1000) AS distance
	FROM places AS p
	WHERE is_deleted = false
	AND (earth_box(ll_to_earth($1, $2), $3) @> ll_to_earth (p.lat, p.lng)
	AND earth_distance(ll_to_earth ($1, $2), ll_to_earth (p.lat, p.lng)) < $3)
	ORDER BY distance, id LIMIT $4 OFFSET $5
	`

	rows, err := config.GetDB().Query(sql, params.Lat, params.Lng, params.Distance*1000, params.Limit, params.Page)
	helpers.CheckAndThrowException(err)
	defer rows.Close()

	var results []*validations.NearbyPlaceResult

	for rows.Next() {
		var p validations.NearbyPlaceResult

		if err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.ResourceId, &p.Lat, &p.Lng, &p.Distance); err != nil {
			log.Fatal(err)
		}

		results = append(results, &p)
	}

	return results
}
