package helper

import (
	"todoapp/internal/database"

	"github.com/google/uuid"
)

type CustomUser struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Password string    `json:"-"`
	Id       uuid.UUID `json:"id"`
}

func CustomUserConvertor(user database.User) CustomUser {
	return CustomUser{
		Name:     user.Name,
		Password: user.Password,
		Age:      int(user.Age),
		Id:       user.ID,
	}
}
