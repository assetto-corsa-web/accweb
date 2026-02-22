package cfg

import (
	"errors"
	"strings"

	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
)

type User struct {
	Username string
	Password string
	Role     string
}

func (u *User) UnmarshalYAML(data []byte) error {
	parts := strings.Split(strings.Trim(string(data), " \r\n\"'"), ":")
	if len(parts) != 3 {
		return errors.New("invalid user format, expected role:username:password")
	}

	u.Username = parts[1]
	u.Password = parts[2]
	u.Role = parts[0]

	return nil
}

func (u User) MarshalYAML() ([]byte, error) {
	return []byte(u.Role + ":" + u.Username + ":" + u.Password), nil
}

type Users []User

func (u *Users) Add(role, username, rawPassword string) error {
	password, err := helper.HashPassword(rawPassword)
	if err != nil {
		return err
	}

	newUser := User{
		Role:     role,
		Username: username,
		Password: password,
	}

	*u = append(*u, newUser)
	return nil
}

func (u Users) Update(username, role string, rawPassword *string) error {
	for i, user := range u {
		if user.Username == username {
			if rawPassword != nil {
				password, err := helper.HashPassword(*rawPassword)
				if err != nil {
					return err
				}
				u[i].Password = password
			}

			u[i].Role = role
			return nil
		}
	}

	return errors.New("user not found")
}

func (u *Users) DeleteByUsername(username string) error {
	for i, user := range *u {
		if user.Username == username {
			*u = append((*u)[:i], (*u)[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found")
}

func (u Users) ResetPassword(username string) (string, error) {
	for i, user := range u {
		if user.Username == username {
			newPassword, err := helper.GenerateRandomPassword()
			if err != nil {
				return "", err
			}

			password, err := helper.HashPassword(newPassword)
			if err != nil {
				return "", err
			}

			u[i].Password = password
			return newPassword, nil
		}
	}

	return "", errors.New("user not found")
}

func (u Users) ValidateUserAndPassword(username, password string) (*User, error) {
	for _, user := range u {
		if user.Username == username {
			if helper.CheckPasswordHash(password, user.Password) {
				return &user, nil
			}

			return nil, errors.New("invalid password")
		}
	}

	return nil, errors.New("user not found")
}
