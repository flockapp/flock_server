package api

import (
	"encoding/json"
	"fmt"
	"github.com/flockapp/flock_server/models"
	"net/http"
)

func V0_Get_API(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, models.Response{
		Success: true,
		Message: "Server is healthy.",
	}, 200)
}

func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	//dj, err := json.MarshalIndent(d, "", "  ")
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	//CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(c)
	fmt.Printf("Response: %v\n", d)
	fmt.Fprintf(w, "%s", dj)
}
