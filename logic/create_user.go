package logic

import (
	"fmt"

	"./../models"
)

func CreateUser(user *models.User, password string) (err error) {
	user.CryptedPassword = password
	user.Id = 12314123123
	fmt.Println(user)
	return
}
