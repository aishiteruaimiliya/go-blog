package blogs

// 用户账号信息
type Account struct {
	Account string `json:"account" gorm:"type:varhcar(50);primary_key"`
	Password string `json:"password"gorm:"type:varchar(50);not null"`
	Banned bool `json:"banned"gorm:"type:bool;"`
	AID string `json:"aid"gorm:"type:char(128);not null"`
}
