package blogs

// 作者，也就是用户
type Author struct {
	Aid string `json:"aid"gorm:"type:char(128);primary_key"`
	Name string `json:"name"gorm:"type:varchar(50)"`
	Gender bool `json:"gender"gorm:"type:bool"`
	Age int `json:"age"gorm:"type:int"`
	Contact string `json:"contact"gorm:"type:char(20)"`
}