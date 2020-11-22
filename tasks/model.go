package tasks

// WallResponseItem vk wall response item
type WallResponseItem struct {
	ID       int    `json:"id"`
	FromID   int    `json:"from_id"`
	OwnerID  int    `json:"owner_id"`
	Date     int64  `json:"date"`
	PostType string `json:"post_type"`
	Text     string `json:"text"`
}

// WallResponseData vk wall response data
type WallResponseData struct {
	Count int                `json:"count"`
	Items []WallResponseItem `json:"items"`
}

// WallResponse vk wall response
type WallResponse struct {
	Response WallResponseData `json:"response"`
}
