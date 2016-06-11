package models

type Guest struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
}

func GetGuestsByEventId(eventId int64) ([]Guest, error) {
	guestList := []Guest{}
	query := db.Raw("SELECT * FROM user INNER JOIN userEvent ON userEvent.eventId = ? WHERE userEvent.userId = user.id", eventId).Scan(&guestList)
	if err := query.Error; err != nil && !query.RecordNotFound() { //Discards case where event is empty
		return nil, err
	}
	return guestList, nil
}