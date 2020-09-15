package blogs

// 作者，也就是用户
type User struct {
	Account   string `json:"account" gorm:"type:varchar(50);primary_key"`
	Password  string `json:"password"gorm:"type:varchar(50);not null"`
	Banned    bool   `json:"banned"gorm:"type:bool;"`
	BanReason string `json:"ban_reason"gorm:"type:varchar(128)"`
	Name      string `json:"name"gorm:"type:varchar(50)"`
	Gender    bool   `json:"gender"gorm:"type:bool"`
	Age       int    `json:"age"gorm:"type:int"`
	Contact   string `json:"contact"gorm:"type:char(20)"`
}
