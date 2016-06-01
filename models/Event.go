package models

import "time"

type Event struct {
	Id     int64     `json:"id"`
	HostId int64     `json:"-"`
	Name   string    `json:"name"`
	Guests []User    `gorm:"many2many:user_events;" json:"guests"`
	Time   time.Time `json:"time"`
	Lat    float64   `json:"lat"`
	Lng    float64   `json:"lng"`
}
