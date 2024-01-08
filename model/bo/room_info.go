package bo

type RoomInfo struct {
	RoomTitle    string `json:"room_title"`
	RoomStatus   string `json:"room_status"`
	AnchorId     int64  `json:"anchor_id"`
	AnchorName   string `json:"anchor_name"`
	AnchorAvatar string `json:"anchor_avatar"`
}
