package models

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Debug   string `json:"debug,omitempty"`
}
