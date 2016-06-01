package models

type User struct {
	Id       int64   `json:"id"`
	Username string  `json:"username"`
	FullName string  `json:"fullName"`
	Password string  `json:"password"`
	Events   []Event `gorm:"many2many:user_events;"`
}
