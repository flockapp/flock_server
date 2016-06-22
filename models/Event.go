package models

import (
	"fmt"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
	"math"
	"sort"
)

type Event struct {
	Id     int64   `json:"id"`
	HostId int64   `json:"-"`
	Name   string  `json:"name"`
	Time   int64   `json:"time"`
	Cost   int64   `json:"cost"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Types  []string `json:"types,omitempty" gorm:"-"`
	Recommendations RecommendationList `json:"-" gorm:"-"`
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

func (e *Event) GetPriceLevels() (maps.PriceLevel, maps.PriceLevel) {
	lower, upper := int(e.Cost) - 1, int(e.Cost) + 1
	if lower < 0 {
		lower = 0
	}
	if upper > 4 {
		upper = 4
	}
	return maps.PriceLevel(string(lower)), maps.PriceLevel(string(upper))
}

func (e *Event) VerifyHostFromRequest(user User) error {
	if user.Id != e.HostId {
		return fmt.Errorf("Event hostId %v does not match user id %v\n", e.HostId, user.Id)
	}
	return nil
}

func (e *Event) GetRecommendations(query string) error {
	client, err := maps.NewClient(maps.WithAPIKey(Conf.GoogleApiKey))
	if err != nil {
		return err
	}
	minCost, maxCost := e.GetPriceLevels()
	searchRequest := maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: e.Lat,
			Lng: e.Lng,
		},
		Radius: 5000, //5km radius
		MinPrice: minCost,
		MaxPrice: maxCost,
		Keyword: query,
	}
	fmt.Printf("debug: query: %v\n minCost: %v\n maxCost: %v\n", query, minCost, maxCost)
	resp, err := client.NearbySearch(context.Background(), &searchRequest)
	if err != nil {
		return err
	}
	searchResults := resp.Results
	scoreMap := make(map[string]float64)
	activities, err := GetActivitiesByEventId(e.Id)
	if err != nil {
		return err
	}
	for _, activity := range activities {
		for _, item := range searchResults {
			for _, typeName := range activity.Types {
				if stringInSlice(typeName, item.Types) {
					scoreMap[item.PlaceID] += 1
				}
			}
			scoreMap[item.PlaceID] -= math.Abs(float64(int(activity.Cost) - item.PriceLevel))
			scoreMap[item.PlaceID] -= math.Abs(activity.Rating - float64(item.Rating))
		}
	}
	recommendations := make(RecommendationList, len(searchResults))
	for i, result := range searchResults {
		recommendations[i] = Recommendation{
			Score: scoreMap[result.PlaceID],
			PlaceId: result.PlaceID,
			Name: result.Name,
			Rating: float64(result.Rating),
			Cost: int64(result.PriceLevel),
			Address: result.FormattedAddress,
			Types: result.Types,
		}
	}
	sort.Sort(recommendations)
	e.Recommendations = recommendations
	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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


