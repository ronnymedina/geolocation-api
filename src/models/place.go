package models

import "time"

type Place struct {
	Id          int64
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ResourceId  int16     `json:"resource_id"`
	Lat         float32   `json:"lat"`
	Lng         float32   `json:"lng"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
