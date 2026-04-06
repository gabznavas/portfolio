package handlers

type CreateLocationRequestBody struct {
	Username  string  `json:"username" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
