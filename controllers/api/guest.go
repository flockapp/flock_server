package api

import (
	"net/http"
	"github.com/flockapp/flock_server/utils"
	"fmt"
	"github.com/flockapp/flock_server/models"
	"strconv"
)

func API_Get_Guests(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetCurrentUser(r)
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		failResp := models.Response{Success: false, Message: "Internal server error.", }
		JSONResponse(w, failResp, 500)
		return
	}
	eventId, err := strconv.Atoi(r.URL.Query().Get("eventId"))
	if err != nil {
		fmt.Printf("Invalid query string: %v\n", err)
		failResp := models.Response{Success: false, Message: "Invalid query string", }
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