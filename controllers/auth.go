package controllers

import (
	"fmt"
	"github.com/flockapp/flock_server/controllers/api"
	"github.com/flockapp/flock_server/models"
	"github.com/flockapp/flock_server/utils"
	"net/http"
	"encoding/json"
)

func AUTH_Post_Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	if err := decoder.Decode(&user); err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal Server Error.",
		}, 500)
		return
	}
	if err := user.Save(); err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal Server Error.",
		}, 500)
		return
	}
	api.JSONResponse(w, models.Response{
		Success: true,
		Message: "Successfully registered user",
	}, 200)
}

func AUTH_Post_Login(w http.ResponseWriter, r *http.Request) {
	//Parse request form
	if err := r.ParseForm(); err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Failed to login",
		}, 500)
		return
	}
	//Extract form values
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := models.FindUserByUsername(username)
	if err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Invalid username/password",
		}, 400)
		return
	}
	//Compare password in request form against db record
	if err := user.VerifyPassword(password); err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Invalid username/password",
		}, 400)
		return
	}
	token, err := utils.CreateToken(user.Id)
	if err != nil {
		fmt.Printf("Error Occured: %v\n", err)
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal server error",
		}, 500)
		return
	}
	api.JSONResponse(w, models.Response{
		Success: true,
		Message: "Successfully logged in",
		Data:    token,
	}, 200)
}
