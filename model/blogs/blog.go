package blogs

import "time"
// 博客具体内容
type Blog struct {
	BID string `json:"id"gorm:"type:char(128);primary_key"`
	Title string `json:"title"gorm:"type:varchar(50)"`
	Author string `json:"author"gorm:"type:varchar(20)"`
	TS time.Time `json:"ts"gorm:"type:timestamp"`
	Content string `json:"content"gorm:"type:text"`
}
