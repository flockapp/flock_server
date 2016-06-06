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
	decoder := json.NewDecoder(r.Body)
	pending_user := models.User{}
	if err := decoder.Decode(&pending_user); err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Internal Server Error.",
		}, 500)
		return
	}
	user, err := models.FindUserByUsername(pending_user.Username)
	if err != nil {
		api.JSONResponse(w, models.Response{
			Success: false,
			Message: "Invalid username/password",
		}, 400)
		return
	}
	//Compare password in request form against db record
	if err := user.VerifyPassword(pending_user.Password); err != nil {
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
