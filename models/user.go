package models

import (
	"errors"
	"fmt"

	"github.com/hainguyen267/go-rest-api/utils"

	"github.com/hainguyen267/go-rest-api/db"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`	
}


func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("error when create the prepare statement")
		return err
	}

	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hashPassword
	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		fmt.Println("error when executing the statement")
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"	

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	correctPassword := utils.ComparePassword(retrievedPassword, u.Password)

	if !correctPassword {
		return errors.New("wrong password")
	}

	return nil
}