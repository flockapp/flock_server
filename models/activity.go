package models

import "time"

type Activity struct {
	Id        int64     `json:"id"`
	EventId   int64     `json:"eventId"`
	PlaceId   string    `json:"placeId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Types     []Type    `json:"types" gorm:"-"`
	Desc      string    `json:"desc"`
}
