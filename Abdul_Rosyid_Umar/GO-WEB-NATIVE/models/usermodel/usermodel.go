package usermodel

import (
	"errors"
	"go-web-native/config"
	"go-web-native/entities"
)

func Authenticate(username, password string) (*entities.User, error) {
	query := "SELECT id, username, password, role FROM users WHERE username = ? AND password = ?"
	row := config.DB.QueryRow(query, username, password)

	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &user, nil
}
