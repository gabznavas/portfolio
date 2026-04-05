package handlers

type ListNearbyByPositionResponse struct {
	Username  string  `json:"username"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
