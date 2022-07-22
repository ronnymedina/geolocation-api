package validations

type CreatePlace struct {
	Name        string  `json:"name" binding:"required,min=5,max=50"`
	Description string  `json:"description" binding:"required,min=3,max=150"`
	ResourceId  int16   `json:"resource_id" binding:"required"`
	Lat         float32 `json:"lat" binding:"required,latitude"`
	Lng         float32 `json:"lng" binding:"required,longitude"`
}

type UpdatePlace struct {
	Name        string `json:"name" binding:"required,min=5,max=50"`
	Description string `json:"description" binding:"required,min=3,max=150"`
	ResourceId  int16  `json:"resource_id" binding:"required"`
}

type NearbyPlaces struct {
	Lat      float32 `json:"lat" binding:"required,latitude"`
	Lng      float32 `json:"lng" binding:"required,longitude"`
	Distance int     `json:"distance" binding:"required"`
	Page     int     `json:"page" binding:"required"`
	Limit    int     `json:"limit" binding:"required"`
}

type NearbyPlaceResult struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ResourceId  int16   `json:"resourceId"`
	Lat         float32 `json:"lat"`
	Lng         float32 `json:"lng"`
	Distance    float32 `json:"distance"`
}
