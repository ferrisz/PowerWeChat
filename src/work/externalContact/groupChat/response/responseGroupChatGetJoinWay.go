package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type JoinWay struct {
	ConfigID       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseID     int      `json:"room_base_id"`
	ChatIDList     []string `json:"chat_id_list"`
	QrCode         int      `json:"qr_code"`
	State          int      `json:"state"`
}

type ResponseGroupChatGetJoinWay struct {
	response.ResponseWork
	JoinWay *JoinWay `json:"join_way"`
}
