package models

type Event struct {
	Id     int64   `json:"id"`
	HostId int64   `json:"-"`
	Name   string  `json:"name"`
	Time   int64   `json:"time"`
	Cost   int64   `json:"cost"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Types  []string `json:"types,omitempty" gorm:"-"`
}

func (e *Event) Save() error {
	for _, name := range e.Types {
		typeInst, err := GetTypeByName(name)
		if err != nil {
			return err
		}
		if err := db.Exec("INSERT INTO `eventType` VALUES (?, ?)", e.Id, typeInst.Id).Error; err != nil {
			return err
		}
	}
	return db.Save(&e).Error
}

func (e *Event) AddGuestById(id int64) error {
	err := db.Exec("INSERT INTO userEvent VALUES (?, ?)", id, e.Id).Error
	return err

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

func GetGuestEventsByUserId(id int64) (*[]Event, error) {
	eventList := []Event{}
	query := db.Raw("SELECT * FROM event JOIN userEvent ON userEvent.userId = ? WHERE event.id = userEvent.eventId", id).Scan(&eventList)
	return &eventList, query.Error
}

func GetEventById(id int64) (*Event, error) {
	event := Event{}
	err := db.Where("id = ?", id).First(&event, Event{}).Error
	return &event, err
}


