package models

import "time"

type Activity struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	EventId   int64     `json:"-"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	ImageUrl  string    `json:"imageUrl"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Type      string    `json:"type"`
	Desc      string    `json:"desc"`
}
