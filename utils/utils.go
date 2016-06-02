package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/flockapp/flock_server/models"
	"net/http"
	"github.com/gorilla/context"
	"errors"
)

func CreateToken(id int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["id"] = id
	tokenString, err := token.SignedString([]byte(models.Conf.SigningKey))
	return tokenString, err
}

func ClearRequest(r *http.Request) {
	context.Clear(r)
}

func SetCurrentUser(r *http.Request, user models.User) {
	context.Set(r, "user", user)
}

func GetCurrentUser(r *http.Request) (models.User, error) {
	user := context.Get(r, "user")
	if user == nil {
		return models.User{}, errors.New("Unable to get user, please check User Id")
	}
	return user.(models.User), nil
}

