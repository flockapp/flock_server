package models

import "fmt"

type Event struct {
	Id     int64   `json:"id"`
	HostId int64   `json:"-"`
	Name   string  `json:"name"`
	Time   int64   `json:"time"`
	Cost   int64   `json:"cost"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Types  []int64 `json:"types" gorm:"-"`
}

func (e *Event) Save() error {
	if err := db.Save(&e).Error; err != nil {
		return err
	}
	for _, val := range e.Types {
		fmt.Println(val)
		if err := db.Exec("INSERT INTO `eventType` VALUES (?, ?)", e.Id, val).Error; err != nil {
			return err
		}
	}
	return nil
}

//TODO: Add event field validation

func GetEventsByUserId(id int64) (*[]Event, error) {
	eventList := []Event{}
	query := db.Where("host_id = ?", id).Find(&eventList, &Event{})
	if query.RecordNotFound() || query.Error == nil {
		return &eventList, nil
	}
	return nil, query.Error
}

func GetEventById(id int64) (*Event, error) {
	event := Event{}
	err := db.Where("id = ?", id).First(&event, Event{}).Error
	return &event, err
}
