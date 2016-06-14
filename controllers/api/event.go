package api

import (
	"encoding/json"
	"fmt"
	"github.com/flockapp/flock_server/models"
	"github.com/flockapp/flock_server/utils"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func API_Get_Events(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Error getting user instance: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Invalid user token",
		}, 400)
		return
	}
	eventList, err := models.GetEventsByUserId(user.Id)
	if err != nil {
		fmt.Printf("Error getting events: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error.",
		}, 500)
		return
	}
	JSONResponse(w, models.Response{
		Success: true,
		Data:    eventList,
		Message: "Successfully retrieved events",
	}, 200)
}

func API_Create_Event(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		fmt.Printf("Failed to parse json: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error.",
		}, 500)
		return
	}
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error.",
		}, 500)
		return
	}
	event.HostId = user.Id
	if err := event.Save(); err != nil {
		fmt.Printf("Failed to get save event: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error.",
		}, 500)
		return
	}
	JSONResponse(w, models.Response{
		Success: true,
		Data:    event,
		Message: "Successfully created event",
	}, 200)
}

func API_Get_Guest_Events(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		failResp := models.Response{Success: false, Message: "Invalid token",}
		JSONResponse(w, failResp, 400)
		return
	}
	eventList, err := models.GetGuestEventsByUserId(user.Id)
	if err != nil {
		fmt.Printf("Unable to get guest events: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error."}
		JSONResponse(w, failResp, 500)
		return
	}
	successResp := models.Response{Success: true, Data: eventList, Message: "Successfully retrieved guest events"}
	JSONResponse(w, successResp, 200)
}

func API_Put_User_Into_Event(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Error getting user instance: %v\n", err)
		failResp := models.Response{Success: false, Message: "Invalid user token", }
		JSONResponse(w, failResp, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	eventContainer := models.Event{}
	if err := decoder.Decode(&eventContainer); err != nil {
		fmt.Printf("Error decoding json to event object: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error."}
		JSONResponse(w, failResp, 500)
		return
	}
	event, err := models.GetEventById(eventContainer.Id)
	if err != nil {
		fmt.Printf("Error retrieving event: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error."}
		JSONResponse(w, failResp, 500)
		return
	}
	if err := event.AddGuestById(user.Id); err != nil {
		fmt.Printf("Error adding guest to event: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error."}
		JSONResponse(w, failResp, 500)
		return
	}
	successResp := models.Response{Success: true, Message: "Successfully added guest to event"}
	JSONResponse(w, successResp, 200)
}


func API_Get_Guests(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error.", }
		JSONResponse(w, failResp, 500)
		return
	}
	eventId, err := strconv.Atoi(mux.Vars(r)["eventId"])
	if err != nil {
		fmt.Printf("Invalid event ID: %v\n", err)
		failResp := models.Response{Success: false, Message: "Invalid event ID", }
		JSONResponse(w, failResp, 400)
		return
	}
	event, err := models.GetEventById(int64(eventId))
	if err != nil {
		fmt.Printf("Invalid event id: %v\n", err)
		failResp := models.Response{Success: false, Message: "Invalid event id", }
		JSONResponse(w, failResp, 400)
		return
	}
	if event.HostId != user.Id {
		failResp := models.Response{Success: false, Message: "User does not own event", }
		JSONResponse(w, failResp, 400)
		return
	}
	guestList, err := models.GetGuestsByEventId(int64(eventId))
	if err != nil {
		fmt.Printf("Could not retrieve guest list: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error.", }
		JSONResponse(w, failResp, 500)
		return
	}
	JSONResponse(w, models.Response{
		Success: true,
		Data: guestList,
		Message: "Successfully retrieve guest list",
	}, 200)
}
