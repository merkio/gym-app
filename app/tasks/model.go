package tasks

// WallResponseItem vk wall response item
type WallResponseItem struct {
	ID          int          `json:"id"`
	FromID      int          `json:"from_id"`
	OwnerID     int          `json:"owner_id"`
	Date        int64        `json:"date"`
	PostType    string       `json:"post_type"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// Attachments vk post item
type Attachment struct {
	Type  string `json:"type"`
	Video Video  `json:"video"`
}

// Video attachment in vk post item
type Video struct {
	ID      int64  `json:"id"`
	OwnerID int64  `json:"owner_id"`
	Title   string `json:"title"`
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
