package blogs

import "time"
// 用户私聊信息
type Message struct {
	SendId string `json:"send_id"gorm:"type:char(128);primary_key"`
	RecvId string `json:"recv_id"gorm:"type:char(128);primary_key"`
	Msg string `json:"msg"gorm:"type:text"`
	TS time.Time `json:"ts"gorm:"type:timestamp"`
}
