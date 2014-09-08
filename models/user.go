package models

type User struct {
	Id		int64
	Email           string
	FirstName       string
	LastName        string
	CryptedPassword string
}

//CREATE TABLE users
//(
//id SERIAL PRIMARY KEY,
//email VARCHAR(255),
//first_name VARCHAR(255),
//last_name VARCHAR(255),
//password VARCHAR(255)
//);
