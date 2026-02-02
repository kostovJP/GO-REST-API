package models

import (
	"example.com/REST-API/db"
	"example.com/REST-API/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	saveQuery := `
	INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(saveQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// we don't want to store the raw passwords in the database...
	// we must first hash them
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	// now we store the hashedPassword instead
	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	//we also want to retrieve the id that was created by autoincrement
	userID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = userID
	return nil
}
