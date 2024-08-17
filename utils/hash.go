package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	fmt.Println("Hashing the password")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	fmt.Println("Done hashing the password")

	return string(bytes), err
}

func ComparePassword(hashedPassword string, plainTextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
	return err == nil
}