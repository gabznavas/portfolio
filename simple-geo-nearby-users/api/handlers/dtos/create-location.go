package handlers

type CreateLocationRequestBody struct {
	Username  string  `json:"username" binding:"required"`
	Latitude  float32 `json:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" binding:"required"`
}
