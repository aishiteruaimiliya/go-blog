package blogs

import "time"
// 对博客的评论
type Comment struct {
	CID string 	`json:"cid"gorm:"type:char(128);primary_key"`
	BID string `json:"bid"gorm:"type:char(128)"`
	AID string `json:"aid"gorm:"type:char(128)"`
	TS time.Time `json:"ts"gorm:"type:timestamp"`
	Msg string `json:"msg"gorm:"type:text"`
}