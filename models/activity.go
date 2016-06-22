package models

import "time"

type Activity struct {
	Id        int64     `json:"id"`
	EventId   int64     `json:"eventId"`
	PlaceId   string    `json:"placeId"`
	Rating    int64 `json:"rating"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Cost      float64 `json:"cost"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Types     []Type    `json:"types" gorm:"-"`
	Desc      string    `json:"desc"`
}

func GetActivitiesByEventId(eventId int64) (*[]Activity, error) {
	activities := []Activity{}
	query := db.Where("eventId = ?", eventId).Find(&activities, &Activity{})
	err := query.Error
	if !query.RecordNotFound() && err != nil {
		return activities, err
	}
	return activities, nil
}

func (a *Activity) Save() error {
	for typeName := range a.Types {
		typeInst, err := GetTypeByName(typeName)
		if err != nil {
			return err
		}
		if err := db.Exec("INSERT INTO `activityType` VALUES (?, ?)", a.Id, typeInst.Id).Error; err != nil {
			return err
		}
	}
	return db.Save(&a).Error
}