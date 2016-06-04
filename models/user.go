package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64   `json:"id"`
	Username string  `json:"username"`
	FullName string  `json:"fullName"`
	Password string  `json:"password"`
	Events   []Event `gorm:"many2many:user_events;"`
}

func FindUserById(id int64) (User, error) {
	user := User{}
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}

func FindUserByUsername(username string) (User, error) {
	user := User{}
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (user *User) Save() error {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passBytes)
	err = db.Save(&user).Error
	return err
}
