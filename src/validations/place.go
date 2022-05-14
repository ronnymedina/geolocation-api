package validations

type CreatePlace struct {
	Name        string  `json:"name" binding:"required,min=5,max=50"`
	Description string  `json:"description" binding:"required,min=3,max=150"`
	Resource_id int16   `json:"resource_id" binding:"required"`
	Lat         float32 `json:"lat" binding:"required,latitude"`
	Lng         float32 `json:"lng" binding:"required,longitude"`
}
