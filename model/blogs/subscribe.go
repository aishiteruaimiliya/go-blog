package blogs

type Subscribe struct {
	Subscribe  string `json:"subscribe"gorm:"type:varchar(128);primary_key"`
	Subscribed string `json:"subscribed"gorm:"type:varchar(128);primary_key"`
}
