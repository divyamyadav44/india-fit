package model

import "time"

type User struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobile_number"`
	DOB          string    `json:"dob"`
	CreatedAt    time.Time `json:"created_at"`
}

type NewUser struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	MobileNumber   string `json:"mobile_number"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password"`
	DOB            string `json:"dob"`
}
