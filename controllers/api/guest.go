package api

import (
	"net/http"
	"github.com/flockapp/flock_server/utils"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/flockapp/flock_server/models"
)

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

func API_Put_Guest_Into_Event(w http.ResponseWriter, r *http.Request) {
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


func API_Get_Guests_From_Event(w http.ResponseWriter, r *http.Request) {
	user, _ := utils.GetCurrentUser(r)
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
	if err := event.VerifyHostFromRequest(user); err != nil {
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
