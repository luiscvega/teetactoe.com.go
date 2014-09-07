package models

type User struct {
	Id		int
	Email           string
	FirstName       string
	LastName        string
	CryptedPassword string
}
