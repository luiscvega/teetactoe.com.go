package logic

import (
	"fmt"

	"./../models"
)

func CreateUser(user *models.User) (err error) {
	fmt.Println("CREATED A USER!", user)

	return
}
