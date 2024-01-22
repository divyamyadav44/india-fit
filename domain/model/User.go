package model

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	DOB          string `json:"dob"`
}

type NewUser struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Password     string `json:"password"`
	DOB          string `json:"dob"`
}
