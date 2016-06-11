package controllers

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/flockapp/flock_server/models"
	"github.com/flockapp/flock_server/controllers/api"
	"github.com/flockapp/flock_server/utils"
)

func RequireUserToken(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			api.JSONResponse(w, models.Response{
				Success: false,
				Message: "No authorization header found",
			}, 400)
			return
		}
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
			}
			return []byte(models.Conf.SigningKey), nil
		})
		if err != nil || !token.Valid {
			fmt.Printf("%v\n", err)
			api.JSONResponse(w, models.Response{
				Success: false,
				Message: "Invalid Token",
			}, 400)
			return
		}
		userId, ok := token.Claims["id"].(float64)
		if !ok {
			api.JSONResponse(w, models.Response{
				Success: false,
				Message: "Invalid Token",
			}, 400)
			return
		}
		newId := int64(userId)
		user, err := models.GetUserById(newId)
		if err != nil {
			fmt.Printf("3")
			api.JSONResponse(w, models.Response{
				Success: false,
				Message: "Invalid Token",
			}, 400)
			return
		}
		utils.SetCurrentUser(r, user)
		handler.ServeHTTP(w, r)
		utils.ClearRequest(r)
	}
}