package api

import (
	"net/http"
	"github.com/flockapp/flock_server/utils"
	"fmt"
	"github.com/flockapp/flock_server/models"
)

func API_Get_Events(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Error getting user instance: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error.",
		}, 500)
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
		Data: eventList,
		Message: "Successfully retrieved events",
	}, 200)
}

//func API_Create_Event(w http.ResponseWriter, r *http.Request) {
//
//}