package tasks

type WallResponseItem struct {
	ID       int    `json:"id"`
	FromID   int    `json:"from_id"`
	OwnerID  int    `json:"owner_id"`
	Date     int64  `json:"date"`
	PostType string `json:"post_type"`
	Text     string `json:"text"`
}

type WallResponseData struct {
	Count int                `json:"count"`
	Items []WallResponseItem `json:"items"`
}

type WallResponse struct {
	Response WallResponseData `json:"response"`
}
