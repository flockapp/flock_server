package api

import (
	"encoding/json"
	"fmt"
	"github.com/flockapp/flock_server/models"
	"github.com/flockapp/flock_server/utils"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
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

func API_Get_Event_Details(w http.ResponseWriter, r *http.Request) {
	eventId, err := strconv.Atoi(mux.Vars(r)["eventId"])
	if err != nil {
		fmt.Printf("Invalid event id: %v\n", err)
		failResp := models.Response{Success: false, Message: "Invalid event Id"}
		JSONResponse(w, failResp, 400)
		return
	}
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Unable to get current user: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error."}
		JSONResponse(w, failResp, 500)
		return
	}
	fmt.Printf("debug eventId: %v\n", eventId);
	event, err := models.GetEventById(int64(eventId))
	if err != nil {
		fmt.Printf("Unable to get event: %v\n", err)
		failResp := models.Response{Success: false, Message: "Unable to get event"}
		JSONResponse(w, failResp, 400)
		return
	}
	if event.HostId != user.Id {
		fmt.Println("User does not own event")
		failResp := models.Response{Success: false, Message: "User does not own event"}
		JSONResponse(w, failResp, 400)
		return
	}
	successResp := models.Response{Success: true, Data: event, Message: "Successfully retrieved event details"}
	JSONResponse(w, successResp, 200)

}

