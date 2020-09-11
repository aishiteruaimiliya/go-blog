package blogs

// 用户联系方式
type Contact struct {
	AID string `json:"aid"gorm:"type:char(128);primary_key"`
	Tel string `json:"tel"gorm:"type:char(20)"`
	Email string `json:"email"gorm:"varchar(50)"`
}