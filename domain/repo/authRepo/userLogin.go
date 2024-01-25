package authRepo

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Login(phoneNumber string, password string) (bytes []byte, err error) {
	bytes, err = HashPassword(password)
	log.Println("hashedPasswordFromDb--->", string(bytes), password)

	return

}

func HashPassword(password string) (bytes []byte, err error) {
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err

}
