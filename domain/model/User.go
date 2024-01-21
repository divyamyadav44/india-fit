package model

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}
