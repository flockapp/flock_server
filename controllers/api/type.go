package api

import (
	"net/http"
	"github.com/flockapp/flock_server/models"
	"fmt"
)

func API_Get_Types(w http.ResponseWriter, r *http.Request) {
	types, err := models.GetTypes()
	if err != nil {
		fmt.Printf("Error retrieving types: %v\n", err)
		JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error.",
		}, 500)
		return
	}
	JSONResponse(w, models.Response{
		Success: true,
		Message: "Successfully retrieved types",
		Data: types,
	}, 200)
}