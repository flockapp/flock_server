package models

import "time"

type Event struct {
	Id     int64     `json:"id"`
	HostId int64     `json:"-"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
	Lat    float64   `json:"lat"`
	Lng    float64   `json:"lng"`
}

//func (e *Event) Save() error {
//
//}




func GetEventsByUserId(id int64) (*[]Event, error) {
	eventList := []Event{}
	err := db.Where("hostId = ?", id).Find(Event{}, &eventList).Error
	return &eventList, err
}
