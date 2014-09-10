package models

type User struct {
	Id              int64
	Email           string
	FirstName       string
	LastName        string
	CryptedPassword string
}

func (u *User) FullName() string {
        return u.FirstName + " " + u.LastName
}

//CREATE TABLE users
//(
//id SERIAL PRIMARY KEY,
//email VARCHAR(255),
//first_name VARCHAR(255),
//last_name VARCHAR(255),
//crypted_password VARCHAR(255)
//);
